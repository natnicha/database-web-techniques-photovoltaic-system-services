package controller

import (
	"net/http"
	"os"
	"photovoltaic-system-services/user/repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type loginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(context *gin.Context) {
	var reqBody loginRequestBody
	err := context.BindJSON(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repositories.GetUserByEmail(reqBody.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	isPasswordCorrect := validatePassword(user.Password, reqBody.Password)
	if !isPasswordCorrect {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect email or password"})
		return
	}

	jwt, err := generateJWT(user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	repositories.CreateLogin(user.Id)
	context.JSON(http.StatusOK, gin.H{"access_token": jwt})
}

func validatePassword(hashPassword []byte, password string) (isCorrect bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

var privateKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func generateJWT(userId int) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}
