package repositories

import "photovoltaic-system-services/db"

type Weather struct {
	Latitude       string  `json:"latitude"`
	Longitude      string  `json:"longitude"`
	Datetime       string  `json:"datetime"`
	AirTemperature float32 `json:"air_temperature"`
	Humidity       int     `json:"humidity"`
}

func InsertWeather(weather Weather) (*Weather, error) {
	result := db.Database.Table("weather").Create(&weather)
	if result.Error != nil {
		return nil, result.Error
	}
	return &weather, nil
}
