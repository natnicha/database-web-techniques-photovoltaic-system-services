package repositories

import (
	"photovoltaic-system-services/db"
	"strings"
)

type Product struct {
	Id                int     `json:"id"`
	ProjectId         int     `json:"project_id"`
	SolarPanelModelId int     `json:"solar_panel_model_id"`
	Orientation       string  `json:"orientation"`
	Inclination       float32 `json:"inclination"`
	Area              float32 `json:"area"`
	Geolocation       string  `json:"geolocation"`
}

type ListRequest struct {
	Filter  string `form:"filter"`
	Limit   int    `form:"limit" binding:"required,min=1"`
	Offset  int    `form:"offset" binding:"required"`
	OrderBy string `form:"order_by"`
	SortBy  string `form:"sort_by"`
}

func CreateProduct(product Product) (*Product, error) {
	result := db.Database.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func CheckExistProduct(productId int) int64 {
	var count int64
	db.Database.Table("products").Select("count(id)").Where("id = ?", productId).Count(&count)
	return count
}

func UpdateProject(productId int, product Product) (*Product, error) {
	product.Id = productId
	tx := db.Database.Model(&product).Where("id = ?", productId).UpdateColumn("update_at", "now()").Save(product)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &product, nil
}

func GetProduct(query ListRequest) (*[]Product, error) {
	tx := db.Database.Model(&Product{})
	if len(query.Filter) > 0 {
		filters := strings.Split(query.Filter, ":")
		tx.Where(filters[0] + " = " + filters[1])
	}
	if query.Offset > 0 {
		tx.Offset(query.Offset)
	}
	if query.Limit > 0 {
		tx.Limit(query.Limit)
	}
	if query.SortBy != "" {
		if query.OrderBy != "" {
			tx.Order(query.SortBy + " " + query.OrderBy)
		} else {
			tx.Order(query.SortBy + " ASC")
		}
	} else {
		tx.Order(" id ASC")
	}
	product := []Product{}
	result := tx.Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}
