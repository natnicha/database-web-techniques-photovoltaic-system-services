package middleware

import (
	"net/http"
	"photovoltaic-system-services/auth/controller"

	"github.com/gin-gonic/gin"
)

type Help interface {
	JWTAuthMiddleware() gin.HandlerFunc
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := controller.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		context.Next()
	}
}
