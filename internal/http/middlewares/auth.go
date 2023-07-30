package middlewares

import (
	"crypto/subtle"
	"errors"
	"net/http"
	"strings"

	"github.com/hatamiarash7/uptime-webhook/configs"
	"github.com/hatamiarash7/uptime-webhook/internal/http/resources"

	"github.com/gin-gonic/gin"
)

// IsAuthenticated is a gin middleware function for checking if the request is authenticated or not
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
		return "", errors.New("[AUTH] The bearer keyword not found")
	}

	t := strings.Replace(k, "bearer ", "", 1)
	t = strings.Replace(t, "Bearer ", "", 1)

	return t, nil
}

// Basic authentication for Prometheus
func BasicAuth(handler http.Handler, username, password, realm string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
