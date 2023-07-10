package alert

import (
	"context"
	"net/http"

	"github.com/hatamiarash7/uptime-webhook/configs"
	"github.com/hatamiarash7/uptime-webhook/internal/models"
)

// Repository is an interface for alert repository
type Repository struct {
	client http.Client
	config configs.Config
}

// NewAlertRepository creates a new alert repository
func NewAlertRepository(c configs.Config) *Repository {
	return &Repository{
		client: http.Client{},
		config: c,
	}
}

// CreateAlert creates an alert
func (r *Repository) CreateAlert(ctx context.Context, alert models.Alert) error {
	if r.config.Notifier.Squadcast.IsEnabled {
		if err := r.CreateSquadcastIncident(alert); err != nil {
			return err
		}
	}

	if r.config.Notifier.Telegram.IsEnabled {
		if err := r.CreateTelegramMessage(alert); err != nil {
			return err
		}
	}

	return nil
}
