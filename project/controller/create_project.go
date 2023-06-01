package controller

import (
	"net/http"
	auth "photovoltaic-system-services/auth/controller"
	"photovoltaic-system-services/project/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
)

type requestBody struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"start_at"`
	IsPrinted   bool      `json:"is_printed" default:"false"`
}

func Create(context *gin.Context) {
	reqBody := new(requestBody)
	defaults.SetDefaults(reqBody)

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = validateStruct(reqBody)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
			return
		}
	}

	user, err := auth.GetCurrentUser(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	project, err := repositories.CreateProject(
		repositories.Projects{
			UserId:      user.Id,
			Name:        reqBody.Name,
			Description: reqBody.Description,
			StartAt:     reqBody.StartAt,
			IsPrinted:   reqBody.IsPrinted,
		},
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": project})
}

func validateStruct(obj interface{}) error {
	v := validator.New()
	return v.Struct(obj)
}
