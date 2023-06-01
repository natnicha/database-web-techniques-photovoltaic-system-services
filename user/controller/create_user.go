package controller

import (
	"net/http"
	"time"

	"photovoltaic-system-services/user/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
	"golang.org/x/crypto/bcrypt"
)

type requestBody struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"  validate:"required"`
	Email     string `json:"email"      validate:"required,email"`
	Password  string `json:"password"   validate:"required"`
	IsActive  bool   `json:"is_active"  default:"true"`
}

func Create(context *gin.Context) {
	reqBody := new(requestBody)
	defaults.SetDefaults(reqBody)
	err := context.BindJSON(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validateStruct(reqBody)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
			return
		}
	}

	existingUser, err := repositories.GetUserByEmail(reqBody.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if existingUser != nil {
		context.JSON(http.StatusConflict, gin.H{"error": "The email is already assigned in the system"})
		return
	}

	user, err := prepareUserInfo(*reqBody)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userResult, err := repositories.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": userResult})
}

func validateStruct(obj interface{}) error {
	v := validator.New()
	return v.Struct(obj)
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
		IsActive:  reqBody.IsActive,
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
