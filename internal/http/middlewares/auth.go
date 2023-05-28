package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/arvancloud/uptime-webhook/configs"
	"github.com/arvancloud/uptime-webhook/internal/http/resources"

	"github.com/gin-gonic/gin"
)

func IsAuthenticated(env configs.Environment, token string) gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(token) == 0 || env.IsTesting() {
			context.Next()
			return
		}

		r, err := getAuthorization(context.GetHeader("Authorization"))

		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, resources.JSON{
				Message: "Unauthorized access, check authorization header",
			})
			return
		}

		if token == r {
			context.Next()
			return
		}

		context.AbortWithStatusJSON(http.StatusUnauthorized, resources.JSON{
			Message: "Unauthorized access, check authorization header",
		})
	}
}

func getAuthorization(k string) (string, error) {
	if !strings.Contains(k, "Bearer ") && !strings.Contains(k, "bearer ") {
		return "", errors.New("bearer keyword not found")
	}

	t := strings.Replace(k, "bearer ", "", 1)
	t = strings.Replace(t, "Bearer ", "", 1)

	return t, nil
}
