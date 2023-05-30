package controller

import (
	"net/http"
	"net/mail"
	"time"

	"photovoltaic-system-services/user/repositories"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type requestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsActive  bool   `json:"is_active"`
}

func Create(context *gin.Context) {
	var reqBody requestBody
	err := context.BindJSON(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	isEmailFormat := validateEmailFormat(reqBody.Email)
	if !isEmailFormat {
		context.JSON(http.StatusBadRequest, gin.H{"error": "The email is invalid format"})
		return
	}

	existingUser, err := repositories.GetUserByEmail(reqBody.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if existingUser != nil {
		context.JSON(http.StatusConflict, gin.H{"error": "The email is already assigned in the system"})
		return
	}

	user, err := prepareUserInfo(reqBody)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	userResult, err := repositories.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": userResult})
}

func validateEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func prepareUserInfo(reqBody requestBody) (user repositories.Users, err error) {
	hashPassword, err := bcryptPassword(reqBody.Password)
	if err != nil {
		return user, err
	}

	user = repositories.Users{
		FirstName: reqBody.FirstName,
		LastName:  reqBody.LastName,
		Email:     reqBody.Email,
		Password:  hashPassword,
		UpdateAt:  time.Now(),
	}
	return user, nil
}

func bcryptPassword(password string) (hashPassword []byte, err error) {
	hashPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashPassword, nil
}
