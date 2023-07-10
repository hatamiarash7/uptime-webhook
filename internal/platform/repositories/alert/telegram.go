package alert

import (
	"encoding/json"
	"net/url"
	"sync"

	"github.com/hatamiarash7/uptime-webhook/internal/models"
	log "github.com/sirupsen/logrus"
)

// CreateTelegramMessage sends a telegram message
func (r *Repository) CreateTelegramMessage(alert models.Alert) error {
	var urls []string
	var wg sync.WaitGroup
	results := make(chan string)

	for _, tag := range alert.Data.Service.Tags {
		team := r.config.Notifier.Telegram.Teams[tag][0]
		params := url.Values{}
		params.Add("chat_id", team.Chat)
		params.Add("parse_mode", "markdownv2")
		if team.Topic != "" {
			params.Add("message_thread_id", team.Topic)
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
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			result, err := sendPOSTRequest(url, body)
			if err != nil {
				log.WithError(err).Error("[TELEGRAM] Error sending request to " + url)
				return
			}
			log.Debugf("[TELEGRAM] Result: %s", result)
			results <- result
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return nil
}
