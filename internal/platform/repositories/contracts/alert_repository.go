package contracts

import (
	"context"

	"github.com/arvancloud/uptime-webhook/internal/models"
)

// AlertRepository is an interface for alert repository
type AlertRepository interface {
	CreateAlert(ctx context.Context, value models.Alert) error
}
