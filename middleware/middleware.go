package middleware

import (
	"net/http"
	"photovoltaic-system-services/auth/controller"
	"photovoltaic-system-services/auth/repositories"

	"github.com/gin-gonic/gin"
)

type Help interface {
	JWTAuthMiddleware() gin.HandlerFunc
	CurrentUser(context *gin.Context) (*repositories.Users, error)
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := controller.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		user, err := CurrentUser(context)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "get user ID failed"})
			context.Abort()
			return
		}
		context.Set("user-id", user.Id)
		context.Next()
	}
}

func CurrentUser(context *gin.Context) (*repositories.Users, error) {
	return controller.GetCurrentUser(context)
}
