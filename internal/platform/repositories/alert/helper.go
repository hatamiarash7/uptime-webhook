package alert

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/hatamiarash7/uptime-webhook/internal/models"
	log "github.com/sirupsen/logrus"
)

func sendPOSTRequest(url string, payload []byte, version string) (string, error) {
	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.WithError(err).Error("[HTTP] Error creating new request")
		return "", err
	}

	// Send the request
	client := http.Client{}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "ArvanCloud-Uptime/"+version)
	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Error("[HTTP] Error sending request")
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Error("[HTTP] Error reading response body")
		return "", err
	}

	if resp.StatusCode != http.StatusAccepted &&
		resp.StatusCode != http.StatusOK &&
		resp.StatusCode != http.StatusCreated {
		return string(body), errors.New("[HTTP] Status code is: " + strconv.Itoa(resp.StatusCode))
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
		"state": getAlertColor(alert.Data.Alert.State),
	}

	var payload models.SquadcastIncident
	if alert.Event == "alert_raised" {
		payload = models.SquadcastIncident{
			Message: "[" + alert.Data.Alert.State +
				"] The \"" + alert.Data.Service.ShortName + "\" is " + status,
			Description: "Your `" + alert.Data.Service.DisplayName +
				"` service is " + status +
				" at *" + alert.Data.Alert.CreatedAt.Format("2006-01-02 15:04:05") + "*\n\n" +
				"**Result:** " + alert.Data.Alert.ShortOutput + "\n",
			Tags:      tags,
			Status:    "trigger",
			EventID:   strconv.Itoa(alert.Data.Device.ID),
			Locations: strings.Join(alert.Data.Locations, ", "),
			Retries:   strconv.Itoa(alert.Data.Service.MspNumRetries),
			Type:      alert.Data.Service.MonitoringServiceType,
		}
	} else {
		payload = models.SquadcastIncident{
			Status:  "resolve",
			EventID: strconv.Itoa(alert.Data.Device.ID),
		}
	}

	return payload
}

func getAlertColor(state string) models.SquadcastTag {
	var color string

	switch state {
	case "OK":
		color = "#00D084"
	case "WARNING":
		color = "#FCB900"
	case "CRITICAL":
		color = "#EB144C"
	case "INFO":
		color = "#0693E3"
	default:
		color = "#ABB8C3"
	}

	return models.SquadcastTag{Color: color, Value: state}
}

func formatTelegramMessage(alert models.Alert) models.TelegramMessage {
	var text string

	if alert.Event == "alert_raised" {
		text = "🔥 *Alert - " + alert.Data.Alert.State + "*\n\n"
		text += "📌 *Source:* Uptime\n\n"
		text += "🏷 *Title:* The \"" + alert.Data.Service.ShortName + "\" is down\n\n"
		text += "📄 *Description:* Your `" + alert.Data.Service.DisplayName +
			"` service is down" +
			" at *" + alert.Data.Alert.CreatedAt.Format("2006-01-02 15:04:05") + "*\n\n" +
			"🔍 *Result:* " + alert.Data.Alert.ShortOutput + "\n"

	} else {
		text = "✅ *Resolved*\n\n"
		text += "📌 *Source:* Uptime\n\n"
		text += "🏷 *Title:* The \"" + alert.Data.Service.ShortName + "\" is up\n\n"
		text += "⏱️ *Time:* " + alert.Data.Date.Format("2006-01-02 15:04:05") + "\n\n"
	}

	payload := models.TelegramMessage{
		Text: escapeMarkdown(text),
	}

	return payload
}

func escapeMarkdown(text string) string {
	markdownChars := []string{"_", "`", "[", "]", "(", ")", "~", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}
	escapedChars := []string{"\\_", "\\`", "\\[", "\\]", "\\(", "\\)", "\\~", "\\>", "\\#", "\\+", "\\-", "\\=", "\\|", "\\{", "\\}", "\\.", "\\!"}

	escapedText := text
	for i := 0; i < len(markdownChars); i++ {
		escapedText = strings.ReplaceAll(escapedText, markdownChars[i], escapedChars[i])
	}

	return escapedText
}
