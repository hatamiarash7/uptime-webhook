package application

import (
	alertActions "github.com/arvancloud/uptime-webhook/internal/http/actions"
	"github.com/arvancloud/uptime-webhook/internal/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (a *App) registerRoutes() {
	a.registerMonitoringRoutes()

	api := a.Router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1Alert := v1.Group("/alert", middlewares.IsAuthenticated(a.configs.App.Env, a.configs.API.AccessToken))
			{
				v1Alert.POST("/", alertActions.CreateAlert(a.Repositories.AlertRepository))
			}
		}
	}
}

func (a *App) registerMonitoringRoutes() {
	if !a.configs.App.Env.IsTesting() {
		a.Router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}
}
