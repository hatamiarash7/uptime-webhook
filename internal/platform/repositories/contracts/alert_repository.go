package contracts

import (
	"context"

	"github.com/hatamiarash7/uptime-webhook/internal/models"
)

// AlertRepository is an interface for alert repository
type AlertRepository interface {
	CreateAlert(ctx context.Context, value models.Alert) error
}
