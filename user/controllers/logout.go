package controller

import (
	"net/http"
	"photovoltaic-system-services/user/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Logout(context *gin.Context) {
	userId, _ := context.Get("user-id")
	jwt, err := generateExpiredJWT(userId.(int))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	host := context.GetHeader("Host")
	userAgent := context.GetHeader("User-Agent")
	userLog := repositories.UserLog{
		Type:      "logout",
		UserId:    userId.(int),
		Host:      host,
		UserAgent: userAgent,
	}
	err = repositories.InsertLoginUserLog(userLog)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"access_token": jwt})
}

func generateExpiredJWT(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Unix(),
	})
	return token.SignedString(privateKey)
}
