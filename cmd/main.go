package main

import (
	"context"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/arvancloud/uptime-webhook/configs"
	"github.com/arvancloud/uptime-webhook/internal/platform/application"
	log "github.com/sirupsen/logrus"
)

var config *configs.Config

func init() {
	log.Info("[SETUP] Loading configs")
	cfg, err := configs.Load("configs/config.yml")
	if err != nil {
		log.WithError(err).Fatal("[SETUP] Failed to load configs")
	}
	config = cfg

	if err = application.SetupLogger(cfg); err != nil {
		log.WithError(err).Fatal("[SETUP] Failed to setup logger")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	app, err := application.NewApplication(ctx, config)
	if err != nil {
		log.WithError(err).Fatal("[SETUP] Could not initialize application")
	}

	app.RunHTTPServer(ctx, wg)

	closeSignal := make(chan os.Signal, 1)
	signal.Notify(closeSignal, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	select {
	case <-closeSignal:
		log.Info("[APP] Terminating by os signal")
	case <-ctx.Done():
		log.Info("[APP] Terminating by context cancellation")
	}

	time.Sleep(time.Duration(1000) * time.Millisecond)
	cancel()
	wg.Wait()

	if err = app.Shutdown(); err != nil {
		log.WithError(err).Panic("[APP] Application shutdown encountered error")
	}
}
