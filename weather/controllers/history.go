package controller

import (
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mcuadros/go-defaults"
)

type requestBody struct {
	Geolocation string `json:"geolocation" validate:"required"`
	StartAt     string `json:"start_at" validate:"required"`
	EndAt       string `json:"end_at" validate:"required"`
}

func History(context *gin.Context) {
	reqBody := new(requestBody)
	defaults.SetDefaults(reqBody)
	err := context.BindJSON(&reqBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)

	// immediatly return 202 Accepted response
	go func() {
		defer wg.Done()
		context.JSON(http.StatusAccepted, nil)
		return
	}()

	// parallel making requests for weather
	go func() {
		defer wg.Done()

		startDateTime, err := time.Parse("2006-01-02 15:04:05+01", reqBody.StartAt)
		endDateTime, err := time.Parse("2006-01-02 15:04:05+01", reqBody.EndAt)
		err = ScrapeWeather(
			[]geolocation{extractLatLong(reqBody.Geolocation)},
			startDateTime.Unix(),
			endDateTime.Unix())
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		return
	}()
	return
}
