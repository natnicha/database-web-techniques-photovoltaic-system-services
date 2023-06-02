package repositories

import "photovoltaic-system-services/db"

type Product struct {
	Id                int     `json:"id"`
	ProjectId         int     `json:"project_id"`
	SolarPanelModelId int     `json:"solar_panel_model_id"`
	Orientation       string  `json:"orientation"`
	Inclination       float32 `json:"inclination"`
	Area              float32 `json:"area"`
	Geolocation       string  `json:"geolocation"`
}

func CreateProduct(product Product) (*Product, error) {
	result := db.Database.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}
