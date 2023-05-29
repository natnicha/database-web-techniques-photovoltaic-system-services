package controller

import (
	"net/http"

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
		context.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	user, err := prepareUserInfo(reqBody)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	userResult, err := repositories.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": userResult})
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
