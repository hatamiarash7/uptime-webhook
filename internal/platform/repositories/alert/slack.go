package alert

import (
	"encoding/json"
	"strings"

	"github.com/hatamiarash7/uptime-webhook/internal/models"
	"github.com/hatamiarash7/uptime-webhook/internal/platform/monitoring"
	log "github.com/sirupsen/logrus"
)

// CreateSlackMessage creates an event in Slack
func (r *Repository) CreateSlackMessage(alert models.Alert) error {
	var urls []string

	for _, tag := range alert.Data.Service.Tags {
		url, ok := r.config.Notifier.Slack.Teams[strings.ToLower(tag)]
		if !ok {
			log.Errorf("[SLACK] Team not found for tag: %s", tag)
			r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncSlackSendFailure)})
			continue
		}

		urls = append(urls, url)
	}

	payload := formatSlackMessage(alert)

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	for _, url := range urls {
		r.CallSlack(url, body)
	}

	return nil
}
