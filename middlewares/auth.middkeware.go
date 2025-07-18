package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func MidJwt(context *gin.Context) {
	if headerAuth := context.Request.Header.Get("Authorization"); !strings.HasPrefix(headerAuth, "Bearer ") {
		context.String(http.StatusUnauthorized, "Unauthorized")
		context.Abort()
	}
}
