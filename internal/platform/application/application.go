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
	// Version is the application version
	Version = "DEV"
)

// App is the main application
type App struct {
	configs configs.Config
	Router  *gin.Engine

	Repositories struct {
		AlertRepository contracts.AlertRepository
	}
}

// Shutdown is used to gracefully shutdown the application
func (a *App) Shutdown() (err error) {
	return nil
}

// RunHTTPServer runs the http server
func (a *App) RunHTTPServer(ctx context.Context, wg *sync.WaitGroup) {
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

		log.Debug("[HTTP] Server successfully closed")
	}()
}

func (a *App) registerRouter() {
	log.Info("[Setup] Register router")

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

// NewApplication creates a new application instance
func NewApplication(_ context.Context, config *configs.Config) (*App, error) {
	log.Info("[Setup] Create new application")
	app := &App{configs: *config}

	app.registerRepositories()
	app.registerRouter()
	app.registerRoutes()

	return app, nil
}
