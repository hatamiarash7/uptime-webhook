package application

import (
	"github.com/hatamiarash7/uptime-webhook/internal/platform/repositories/alert"
	log "github.com/sirupsen/logrus"
)

func (a *App) registerRepositories() {
	log.Info("[SETUP] Register repositories")
	a.registerAlertRepository()
}

func (a *App) registerAlertRepository() {
	a.Repositories.AlertRepository = alert.NewAlertRepository(a.configs, a.WorkerPools.AlertPool)
}
