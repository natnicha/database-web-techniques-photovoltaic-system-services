package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"photovoltaic-system-services/weather/repositories"
	"regexp"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseProducts struct {
	Data []repositories.Product `json:"data"`
}

type geolocation struct {
	latitude  string
	longitude string
}

type openWeatherParams struct {
	latitude  string
	longitude string
	start     int64
	end       int64
}

type openWeatherResponse struct {
	Message  string    `json:"message"`
	Code     string    `json:"cod"`
	CityId   int       `json:"city_id"`
	Calctime float32   `json:"calctime"`
	Cnt      int       `json:"cnt"`
	List     []weather `json:"list"`
}

type weather struct {
	Dt   int  `json:"dt"`
	Main main `json:"main"`
}

type main struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	Pressure  float32 `json:"pressure"`
	Humidity  float32 `json:"humidity"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
}

func Daily(context *gin.Context) {
	runtime.GOMAXPROCS(2)

	// immediatly return 202 Accepted response
	go func() {
		context.JSON(http.StatusAccepted, nil)
		return
	}()

	// parallel making requests for weather
	go func() {
		yesterday := time.Now().AddDate(0, 0, -1)
		startDateTime := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, time.Local).Unix()
		endDateTime := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 23, 59, 0, 0, time.Local).Unix()
		products, err := getAllProducts()
		if err != nil {
			log.Println(err.Error())
			return
		}

		uniqueGeolocation := getUniqueGeolocation(products.Data)
		err = ScrapeWeather(uniqueGeolocation, startDateTime, endDateTime)
		if err != nil {
			log.Println(err.Error())
			return
		}
		return
	}()
}

func getAllProducts() (ResponseProducts, error) {
	client := &http.Client{}
	url := "http://localhost:" + os.Getenv("SERVICE_PORT") + "/api/v1/product"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseProducts{}, err
	}

	request.Header.Set("api-key", os.Getenv("APP_API_KEY"))
	resp, err := client.Do(request)
	if err != nil {
		return ResponseProducts{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result ResponseProducts
	err = json.Unmarshal(body, &result)
	if err != nil {
		return ResponseProducts{}, err
	}
	return result, nil
}

func extractLatLong(point string) geolocation {
	regEx := `(?P<lat>[-\d.]+),(?P<long>[-\d.]+)`
	var compRegEx = regexp.MustCompile(regEx)
	matches := compRegEx.FindStringSubmatch(point)

	return geolocation{latitude: matches[1], longitude: matches[2]}
}

func callOpenWeatherAPI(openWeatherParams openWeatherParams) (openWeatherResponse, error) {
	client := &http.Client{}
	url := os.Getenv("OPEN_WEATHER_HISTORY_URL")
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return openWeatherResponse{}, err
	}
	q := request.URL.Query()
	q.Add("lat", openWeatherParams.latitude)
	q.Add("lon", openWeatherParams.longitude)
	q.Add("start", fmt.Sprint(openWeatherParams.start))
	q.Add("end", fmt.Sprint(openWeatherParams.end))
	q.Add("type", os.Getenv("OPEN_WEATHER_DURATION"))
	q.Add("units", os.Getenv("OPEN_WEATHER_UNITS"))
	q.Add("appid", os.Getenv("OPEN_WEATHER_API_KEY"))
	request.URL.RawQuery = q.Encode()

	resp, err := client.Do(request)
	if err != nil {
		return openWeatherResponse{}, err
	}
	if resp.StatusCode != 200 {
		return openWeatherResponse{}, errors.New("Requesting to Open Weather failed")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result openWeatherResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return openWeatherResponse{}, err
	}
	return result, nil
}

func getUniqueGeolocation(products []repositories.Product) []geolocation {
	var uniqueGeolocation []geolocation
	for _, prod := range products {
		geolocation := extractLatLong(prod.Geolocation)
		if len(uniqueGeolocation) == 0 {
			uniqueGeolocation = append(uniqueGeolocation, geolocation)
		} else {
			isExist := false
			for _, geo := range uniqueGeolocation {
				if geolocation.latitude == geo.latitude && geolocation.longitude == geo.longitude {
					isExist = true
				}
			}
			if !isExist {
				uniqueGeolocation = append(uniqueGeolocation, geolocation)
			}
		}
	}
	return uniqueGeolocation
}

func ScrapeWeather(geolocation []geolocation, startDateTime int64, endDateTime int64) error {
	for _, val := range geolocation {
		openWeatherResponse, err := callOpenWeatherAPI(
			openWeatherParams{
				val.latitude,
				val.longitude,
				startDateTime,
				endDateTime,
			},
		)
		if err != nil {
			return err
		}
		for _, weather := range openWeatherResponse.List {
			repositories.InsertWeather(repositories.Weather{
				Latitude:       val.latitude,
				Longitude:      val.longitude,
				Datetime:       time.Unix(int64(weather.Dt), 0).Format(time.RFC3339),
				AirTemperature: weather.Main.Temp,
				Humidity:       int(weather.Main.Humidity),
			})
		}
	}
	return nil
}
