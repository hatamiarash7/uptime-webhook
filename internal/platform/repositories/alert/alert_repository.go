package alert

import (
	"context"
	"net/http"

	"github.com/arvancloud/uptime-webhook/configs"
	"github.com/arvancloud/uptime-webhook/internal/models"
)

// AlertRepository is an interface for alert repository
type AlertRepository struct {
	client http.Client
	config configs.Config
}

// NewAlertRepository creates a new alert repository
func NewAlertRepository(c configs.Config) *AlertRepository {
	return &AlertRepository{
		client: http.Client{},
		config: c,
	}
}

// CreateAlert creates an alert
func (r *AlertRepository) CreateAlert(ctx context.Context, alert models.Alert) error {
	if r.config.Notifier.Squadcast.IsEnabled {
		return r.CreateSquadcastIncident(alert)
	}

	return nil
}
