package alert

import (
	"encoding/json"
	"strings"

	"github.com/hatamiarash7/uptime-webhook/internal/models"
	"github.com/hatamiarash7/uptime-webhook/internal/platform/monitoring"
	log "github.com/sirupsen/logrus"
)

// CreateCustomMessage creates an event in Custom
func (r *Repository) CreateCustomMessage(alert models.Alert) error {
	var urls []string

	for _, tag := range alert.Data.Service.Tags {
		url, ok := r.config.Notifier.Custom.Teams[strings.ToLower(tag)]
		if !ok {
			log.Errorf("[CUSTOM] Team not found for tag: %s", tag)
			r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncCustomSendFailure)})
			continue
		}

		urls = append(urls, url)
	}

	payload := formatCustomMessage(alert)

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	for _, url := range urls {
		r.CallCustom(url, body)
	}

	return nil
}
