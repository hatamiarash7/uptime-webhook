package application

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hatamiarash7/uptime-webhook/configs"
	"github.com/hatamiarash7/uptime-webhook/internal/platform/monitoring"
	"github.com/hatamiarash7/uptime-webhook/internal/platform/repositories/contracts"
	"github.com/panjf2000/ants/v2"

	log "github.com/sirupsen/logrus"
)

// App is the main application
type App struct {
	configs configs.Config
	Router  *gin.Engine

	Repositories struct {
		AlertRepository contracts.AlertRepository
	}

	WorkerPools struct {
		AlertPool *ants.Pool
	}

	Version string

	Monitoring monitoring.Monitor
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
				log.WithError(err).Fatal("[HTTP] " + err.Error())
			}
		}()

		<-ctx.Done()

		shutdownCTX, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		if err := server.Shutdown(shutdownCTX); err != nil {
			log.WithContext(ctx).WithError(err).Error("[HTTP] Could not gracefully shutdown the http server")
		}

		log.Debug("[HTTP] Server successfully closed")
	}()
}

func (a *App) registerRouter() {
	log.Info("[SETUP] Register router")

	switch a.configs.App.Env {
	case configs.Testing:
		gin.SetMode(gin.TestMode)
	case configs.Local, configs.Staging:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	a.Router = gin.New()
	a.Router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/metrics"),
		gin.Recovery(),
	)
}

func (a *App) registerMonitoring() {
	if !a.configs.App.Env.IsTesting() {
		a.Monitoring = monitoring.NewPrometheusMonitor()
		return
	}

	a.Monitoring = monitoring.NewMockMonitor()
}

// NewApplication creates a new application instance
func NewApplication(_ context.Context, config *configs.Config) (*App, error) {
	log.Info("[SETUP] Create new application")
	app := &App{
		configs: *config,
		Version: *&config.Version,
	}

	if err := app.registerAlertPool(); err != nil {
		return nil, err
	}

	app.registerMonitoring()
	app.registerRepositories()
	app.registerRouter()
	app.registerRoutes()

	return app, nil
}
