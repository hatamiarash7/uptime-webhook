package action

import (
	"net/http"

	"github.com/arvancloud/uptime-webhook/internal/http/requests"
	alert_requests "github.com/arvancloud/uptime-webhook/internal/http/requests/alert"
	"github.com/arvancloud/uptime-webhook/internal/http/resources"
	"github.com/arvancloud/uptime-webhook/internal/platform/repositories/contracts"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

// CreateAlert is a gin handler function for creating an alert
func CreateAlert(repository contracts.AlertRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c alert_requests.CreateAlertRequest
		if err := ctx.ShouldBindJSON(&c); err != nil {
			log.WithError(err).Error("[HTTP] Could not bind request body")
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, resources.JSON{
				Errors: requests.GetErrorMessages(err.(validator.ValidationErrors)),
			})
			return
		}

		if err := repository.CreateAlert(ctx, transformRequestToValue(c)); err != nil {
			log.WithError(err).Error("[HTTP] Could not create alert")
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, resources.JSON{
				Message: "could not create alert",
			})
			return
		}

		ctx.JSON(http.StatusOK, resources.JSON{Message: "Created successfully"})
	}
}
