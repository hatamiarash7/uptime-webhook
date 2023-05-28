package alert

import (
	"context"
	"net/http"

	"github.com/arvancloud/uptime-webhook/configs"
	"github.com/arvancloud/uptime-webhook/internal/models"
)

type AlertRepository struct {
	client http.Client
	config configs.Config
}

func NewAlertRepository(c configs.Config) *AlertRepository {
	return &AlertRepository{
		client: http.Client{},
		config: c,
	}
}

func (r *AlertRepository) CreateAlert(ctx context.Context, alert models.Alert) error {
	if r.config.Notifier.Squadcast.IsEnabled {
		return r.CreateSquadcastIncident(alert)
	}

	return nil
}
