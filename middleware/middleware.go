package middleware

import (
	"net/http"
	"os"
	controller "photovoltaic-system-services/auth/controllers"
	"photovoltaic-system-services/auth/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Help interface {
	JWTAuthMiddleware() gin.HandlerFunc
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.GetHeader("api-key") == os.Getenv("APP_API_KEY") {
			userId, _ := strconv.Atoi(context.GetHeader("user-id"))
			context.Set("user-id", userId)
		} else {
			err := controller.ValidateJWT(context)
			if err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				context.Abort()
				return
			}
			user, err := currentUser(context)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": "get user ID failed"})
				context.Abort()
				return
			}
			if user.Id == 0 {
				context.JSON(http.StatusInternalServerError, gin.H{"error": "no user account or an user account has been deleted"})
				context.Abort()
				return
			}
			context.Set("user-id", user.Id)
			context.Set("authorization", context.GetHeader("Authorization"))
		}

		context.Next()
	}
}

func currentUser(context *gin.Context) (*repositories.Users, error) {
	return controller.GetCurrentUser(context)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, HEAD, TRACE, DELETE, PATCH, COPY, LINK, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
