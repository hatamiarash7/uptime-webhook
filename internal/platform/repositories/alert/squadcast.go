package alert

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"

	"github.com/arvancloud/uptime-webhook/internal/models"
	log "github.com/sirupsen/logrus"
)

func (r *AlertRepository) CreateSquadcastIncident(alert models.Alert) error {
	var urls []string
	var wg sync.WaitGroup
	results := make(chan string)

	for _, tag := range alert.Data.Service.Tags {
		urls = append(urls, r.config.Notifier.Squadcast.Teams[strings.ToLower(tag)])
	}

	var status string
	if alert.Data.Alert.IsUp == false {
		status = "down"
	} else {
		status = "up"
	}

	tags := map[string]models.SquadcastTag{
		"state":     {Color: "#d6911a", Value: alert.Data.Alert.State},
		"locations": {Color: "#1bab5c", Value: strings.Join(alert.Data.Locations, ", ")},
	}

	var payload models.SquadcastIncident
	if alert.Event == "alert_raised" {
		payload = models.SquadcastIncident{
			Message:     "The " + alert.Data.Device.DisplayName + " is " + status,
			Description: "Your " + alert.Data.Service.Name + " service is " + status + " at " + alert.Data.Alert.CreatedAt.Format("2006-01-02 15:04:05"),
			Tags:        tags,
			Status:      "trigger",
			EventID:     strconv.Itoa(alert.Data.Alert.ID),
		}
	} else {
		payload = models.SquadcastIncident{
			Status:  "resolve",
			EventID: strconv.Itoa(alert.Data.Alert.ID),
		}
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			result, err := sendPOSTRequest(url, body)
			if err != nil {
				log.WithError(err).Error("Error sending request to " + url)
				return
			}
			results <- result
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return nil
}
