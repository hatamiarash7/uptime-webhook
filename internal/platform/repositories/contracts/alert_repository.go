package contracts

import (
	"context"

	"github.com/arvancloud/uptime-webhook/internal/models"
)

type AlertRepository interface {
	CreateAlert(ctx context.Context, value models.Alert) error
}
