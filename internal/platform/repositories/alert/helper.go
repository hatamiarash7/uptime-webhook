package alert

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/arvancloud/uptime-webhook/internal/models"
)

func sendPOSTRequest(url string, payload []byte) (string, error) {
	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	// Send the request
	client := http.Client{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func formatSquadcastMessage(alert models.Alert) models.SquadcastIncident {
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
			Message: "The " + alert.Data.Service.ShortName + " is " + status,
			Description: "Your `" + alert.Data.Service.DisplayName +
				"` service is " + status +
				" at *" + alert.Data.Alert.CreatedAt.Format("2006-01-02 15:04:05") + "*\n\n" +
				"**State:** " + alert.Data.Alert.State + "\n" +
				"**Output:** " + alert.Data.Alert.ShortOutput + "\n" +
				"**Retries:** " + strconv.Itoa(alert.Data.Service.MspNumRetries),
			Tags:    tags,
			Status:  "trigger",
			EventID: strconv.Itoa(alert.Data.Alert.ID),
		}
	} else {
		payload = models.SquadcastIncident{
			Status:  "resolve",
			EventID: strconv.Itoa(alert.Data.Alert.ID),
		}
	}

	return payload
}
