package alert

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/hatamiarash7/uptime-webhook/internal/models"
	"github.com/hatamiarash7/uptime-webhook/internal/platform/monitoring"
	log "github.com/sirupsen/logrus"
)

// shouldDropAlert checks if the alert matches any of the drop rules for the team.
func shouldDropAlert(alert models.Alert, dropRules []string) bool {
	for _, rule := range dropRules {
		if alert.Data.Alert.Output == rule || alert.Data.Alert.ShortOutput == rule || alert.Event == rule {
			return true
		}
	}
	return false
}

// CreateTelegramMessage sends a telegram message
func (r *Repository) CreateTelegramMessage(alert models.Alert) error {
	var urls []string

	for _, tag := range alert.Data.Service.Tags {
		team, ok := r.config.Notifier.Telegram.Teams[strings.ToLower(tag)]
		if !ok {
			log.Warnf("[Telegram] Team not found for tag: %s", tag)
			r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncTelegramSendFailure)})
			continue
		}

		// Create a list of URLs for every target in the team.
		// For example, a team needs to send a message to 2 chats, so we create 2 URLs.
		for _, t := range team {

			// Check if the alert matches any drop rules for this team.
			if shouldDropAlert(alert, t.DropRules) {
				log.Infof("[Telegram] Alert dropped for team %s due to drop rules: %v", tag, t.DropRules)
				continue
			}

			params := url.Values{}
			params.Add("chat_id", t.Chat)
			params.Add("parse_mode", "markdownv2")
			if t.Topic != "" {
				params.Add("message_thread_id", t.Topic)
			}
			u := r.config.Notifier.Telegram.Host + r.config.Notifier.Telegram.Token + "/sendMessage?" + params.Encode()
			urls = append(urls, u)
		}
	}

	payload := formatTelegramMessage(alert)

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	for _, u := range urls {
		err = r.CallTelegram(u, body)
		if err != nil {
			log.Errorf("[Telegram] Failed to send Telegram message: %s", err)
			r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncTelegramSendFailure)})
		}
	}

	return nil
}
