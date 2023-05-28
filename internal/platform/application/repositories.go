package application

import (
	"github.com/arvancloud/uptime-webhook/internal/platform/repositories/alert"
)

func (a *App) registerRepositories() {
	a.registerAlertRepository()
}

func (a *App) registerAlertRepository() {
	a.Repositories.AlertRepository = alert.NewAlertRepository(a.configs)
}
