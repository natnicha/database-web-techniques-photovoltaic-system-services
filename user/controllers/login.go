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

	if user == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "no user account or an user account has been deleted"})
		return
	}

	isPasswordCorrect := validatePassword(user.Password, reqBody.Password)
	if !isPasswordCorrect {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect email or password"})
		return
	}

	jwt, err := generateValidJWT(user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	host := context.GetHeader("Host")
	userAgent := context.GetHeader("User-Agent")
	userLog := repositories.UserLog{
		Type:      "login",
		UserId:    user.Id,
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

func validatePassword(hashPassword []byte, password string) (isCorrect bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

var privateKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func generateValidJWT(userId int) (string, error) {
	tokenDuration, _ := strconv.Atoi(os.Getenv("JWT_DURATION_MINUTE"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * time.Duration(tokenDuration)).Unix(),
	})
	return token.SignedString(privateKey)
}
