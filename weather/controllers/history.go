package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type requestBody struct {
	Geolocation string `json:"geolocation" validate:"required"`
	StartAt     string `json:"start_at" validate:"required"`
	EndAt       string `json:"end_at" validate:"required"`
}

func History(context *gin.Context) {
	reqBody := new(requestBody)
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

	// immediatly return 202 Accepted response
	go func() {
		context.JSON(http.StatusAccepted, nil)
		return
	}()

	// parallel making requests for weather
	go func() {
		startDateTime, err := time.Parse("2006-01-02 15:04:05-07", reqBody.StartAt)
		endDateTime, err := time.Parse("2006-01-02 15:04:05-07", reqBody.EndAt)
		startIntervalDateTime := startDateTime
		endIntervalDateTime := startDateTime
		for endIntervalDateTime.Before(endDateTime) {
			if endIntervalDateTime.Add(7 * 24 * time.Hour).After(endDateTime) {
				endIntervalDateTime = endDateTime
			} else {
				endIntervalDateTime = endIntervalDateTime.Add(7 * 24 * time.Hour)
			}
			err = ScrapeWeather(
				[]geolocation{extractLatLong(reqBody.Geolocation)},
				startIntervalDateTime.Unix(),
				endIntervalDateTime.Unix())
			if err != nil {
				log.Println(err.Error())
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			startIntervalDateTime = endIntervalDateTime
		}
		return
	}()
	return
}

func validateStruct(obj interface{}) error {
	v := validator.New()
	return v.Struct(obj)
}
