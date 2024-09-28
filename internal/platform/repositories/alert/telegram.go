package alert

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/hatamiarash7/uptime-webhook/internal/models"
	"github.com/hatamiarash7/uptime-webhook/internal/platform/monitoring"
	log "github.com/sirupsen/logrus"
)

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

		// Create a list of URLs for every targets in the team.
		// For example, a team needs to send a message to 2 chats, so we create 2 URLs.
		for _, t := range team {
			params := url.Values{}
			params.Add("chat_id", t.Chat)
			params.Add("parse_mode", "markdownv2")
			// Don't send message if the the topic is matched to any dropped rules.
			if len(r.config.Notifier.Telegram.Drop) > 0 {
				for _, rule := range r.config.Notifier.Telegram.Drop {
					if rule != t.Topic && t.Topic != "" {
						params.Add("message_thread_id", t.Topic)
					}
				}
			} else {
				if t.Topic != "" {
					params.Add("message_thread_id", t.Topic)
				}
			}
			url := r.config.Notifier.Telegram.Host + r.config.Notifier.Telegram.Token + "/sendMessage?" + params.Encode()
			urls = append(urls, url)
		}
	}

	payload := formatTelegramMessage(alert)

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	for _, url := range urls {
		r.CallTelegram(url, body)
	}

	return nil
}
