package application

import (
	"runtime/debug"

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

	return nil
}
