package controller

import (
	"errors"
	"fmt"
	"os"
	"photovoltaic-system-services/auth/repositories"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return errors.New("error getting token or no token provided")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return errors.New("invalid token provided")
	}
	now := time.Now().Unix()
	isExpire := claims.VerifyExpiresAt(now, true)
	if !isExpire {
		return errors.New("token expired")
	}
	return nil
}

func isJWTExpired(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func GetCurrentUser(context *gin.Context) (*repositories.Users, error) {
	err := ValidateJWT(context)
	if err != nil {
		return nil, err
	}
	token, _ := getToken(context)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := int(claims["id"].(float64))

	user, err := repositories.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

var privateKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
