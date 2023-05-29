package application

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/arvancloud/uptime-webhook/configs"
	"github.com/arvancloud/uptime-webhook/internal/platform/repositories/contracts"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

var (
	Version = "DEV"
)

type App struct {
	configs configs.Config
	Router  *gin.Engine

	Repositories struct {
		AlertRepository contracts.AlertRepository
	}
}

func (a *App) Shutdown() (err error) {
	return nil
}

func (a *App) RunHttpServer(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		server := &http.Server{
			Addr:    a.configs.API.ServeAddress,
			Handler: a.Router,
		}

		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.WithError(err).Fatal(err.Error())
			}
		}()

		<-ctx.Done()

		shutdownCTX, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		if err := server.Shutdown(shutdownCTX); err != nil {
			log.WithContext(ctx).WithError(err).Error("could not gracefully shutdown the http server")
		}

		log.Debug("http server successfully closed")
	}()
}

func (a *App) registerRouter() {
	switch a.configs.App.Env {
	case configs.Testing:
		gin.SetMode(gin.TestMode)
	case configs.Local, configs.Staging:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	a.Router = gin.Default()
}

func NewApplication(_ context.Context, config *configs.Config) (*App, error) {
	app := &App{configs: *config}

	app.registerRepositories()
	app.registerRouter()
	app.registerRoutes()

	return app, nil
}