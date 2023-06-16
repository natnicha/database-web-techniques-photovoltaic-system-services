package repositories

import "photovoltaic-system-services/db"

type Weather struct {
	Geolocation    string  `json:"geolocation"`
	Datetime       string  `json:"datetime"`
	AirTemperature float32 `json:"air_temperature"`
	Humidity       int     `json:"humidity"`
}

func InseartWeather(weather Weather) (*Weather, error) {
	result := db.Database.Table("weather").Create(&weather)
	if result.Error != nil {
		return nil, result.Error
	}
	return &weather, nil
}
