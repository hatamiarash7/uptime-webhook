package alert

import (
	"encoding/json"
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

	payload := formatSquadcastMessage(alert)

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
			log.Infof("Result: %s", result)
			results <- result
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return nil
}
