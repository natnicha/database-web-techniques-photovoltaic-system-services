package middleware

import (
	"net/http"
	controller "photovoltaic-system-services/auth/controllers"
	"photovoltaic-system-services/auth/repositories"

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
		context.Request.Header.Add("Access-Control-Allow-Origin", "*")
		context.Request.Header.Add("Access-Control-Allow-Credentials", "true")
		context.Request.Header.Add("Access-Control-Allow-Methods", "GET, PUT, POST, HEAD, TRACE, DELETE, PATCH, COPY, HEAD, LINK, OPTIONS")

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
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
