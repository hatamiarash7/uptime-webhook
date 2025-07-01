package alert

import (
	"encoding/json"
	"strings"

	"github.com/hatamiarash7/uptime-webhook/internal/models"
	"github.com/hatamiarash7/uptime-webhook/internal/platform/monitoring"
	log "github.com/sirupsen/logrus"
)

// CreateMattermostMessage creates an incident in mattermost
func (r *Repository) CreateMattermostMessage(alert models.Alert) error {
	var urls []string

	for _, tag := range alert.Data.Service.Tags {
		url, ok := r.config.Notifier.Mattermost.Teams[strings.ToLower(tag)]
		if !ok {
			log.Warnf("[MATTERMOST] Team not found for tag: %s", tag)
			r.monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.IncMattermostSendFailure)})
			continue
		}

		urls = append(urls, url)
	}

	payload := formatMattermostMessage(alert)

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	for _, url := range urls {
		r.CallMattermost(url, body)
	}

	return nil
}
