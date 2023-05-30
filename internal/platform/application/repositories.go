package application

import (
	"github.com/arvancloud/uptime-webhook/internal/platform/repositories/alert"
	log "github.com/sirupsen/logrus"
)

func (a *App) registerRepositories() {
	log.Info("[Setup] Register repositories")
	a.registerAlertRepository()
}

func (a *App) registerAlertRepository() {
	a.Repositories.AlertRepository = alert.NewAlertRepository(a.configs)
}
