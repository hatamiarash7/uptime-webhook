package alert

import (
	"context"
	"net/http"

	"github.com/hatamiarash7/uptime-webhook/configs"
	"github.com/hatamiarash7/uptime-webhook/internal/models"
	"github.com/panjf2000/ants/v2"
)

// Repository is an interface for alert repository
type Repository struct {
	client  http.Client
	config  configs.Config
	pool    *ants.Pool
	version string
}

// NewAlertRepository creates a new alert repository
func NewAlertRepository(c configs.Config, pool *ants.Pool, version string) *Repository {
	return &Repository{
		client:  http.Client{},
		config:  c,
		pool:    pool,
		version: version,
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
