package alert

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/hatamiarash7/uptime-webhook/internal/models"
	log "github.com/sirupsen/logrus"
)

// CreateTelegramMessage sends a telegram message
func (r *Repository) CreateTelegramMessage(alert models.Alert) error {
	var urls []string

	for _, tag := range alert.Data.Service.Tags {
		team, ok := r.config.Notifier.Telegram.Teams[strings.ToLower(tag)]
		if !ok {
			log.Errorf("[Telegram] Team not found for tag: %s", tag)
			continue
		}

		params := url.Values{}
		params.Add("chat_id", team[0].Chat)
		params.Add("parse_mode", "markdownv2")
		if team[0].Topic != "" {
			params.Add("message_thread_id", team[0].Topic)
		}
		url := r.config.Notifier.Telegram.Host + r.config.Notifier.Telegram.Token + "/sendMessage?" + params.Encode()
		urls = append(urls, url)
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
