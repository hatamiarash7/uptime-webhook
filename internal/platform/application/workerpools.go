package application

import (
	"runtime/debug"

	"github.com/hatamiarash7/uptime-webhook/internal/platform/monitoring"
	"github.com/panjf2000/ants/v2"
	log "github.com/sirupsen/logrus"
)

func (a *App) registerAlertPool() error {
	pool, err := ants.NewPool(a.configs.App.PoolSize, ants.WithPanicHandler(func(err interface{}) {
		log.WithError(err.(error)).WithField("stacktrace", string(debug.Stack())).Error("[POOL] Panic occurred in AlertPool")
	}), ants.WithNonblocking(true))
	if err != nil {
		return err
	}

	a.WorkerPools.AlertPool = pool

	a.Monitoring.Record([]monitoring.Event{monitoring.NewEvent(monitoring.SetAlertPoolCapacity, a.configs.App.PoolSize)})

	a.monitorAlertWorkerPool()

	return nil
}

func (a *App) monitorAlertWorkerPool() {
	_, err := a.scheduler.Every(10).Seconds().Do(func() {
		a.Monitoring.Record(
			[]monitoring.Event{
				monitoring.NewEvent(monitoring.SetActiveJobsInAlertPool,
					a.WorkerPools.AlertPool.Running()),
			},
		)

	})

	if err != nil {
		log.WithError(err).Error("[MONITORING] Error occurred while registering alert pool active jobs metric")
	}
}
