package repositories

import (
	"photovoltaic-system-services/db"
	"strings"
)

type Product struct {
	Id                int     `json:"id"`
	ProjectId         int     `json:"project_id"`
	SolarPanelModelId int     `json:"solar_panel_model_id"`
	Orientation       float32 `json:"orientation"`
	Inclination       float32 `json:"inclination"`
	Area              float32 `json:"area"`
	Geolocation       string  `json:"geolocation"`
	GeneratedEnergy   float32 `json:"generated_energy"`
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

func DeleteProductById(id int) (err error) {
	result := db.Database.Where("id = ?", id).Delete(Product{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetProduct(query ListRequest) (*[]Product, error) {
	tx := db.Database.Model(&Product{})
	if len(query.Filter) > 0 {
		if strings.Contains(query.Filter, ",") {
			filters := strings.Split(query.Filter, ",")
			where := ""
			for _, filter := range filters {
				if where != "" {
					where += " and "
				}
				f := strings.Split(filter, ":")
				where += "cast( " + f[0] + " as varchar)  like '%" + f[1] + "%'"
			}
			tx.Where(where)
		} else {
			filters := strings.Split(query.Filter, ":")
			tx.Where("cast( " + filters[0] + " as varchar)  like '%" + filters[1] + "%'")
		}
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

func UpdateProduct(productId int, product Product) (*Product, error) {
	product.Id = productId
	tx := db.Database.Model(&product).Where("id = ?", productId).Save(product)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = db.Database.Model(&product).Where("id = ?", productId).UpdateColumn("update_at", "now()")
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &product, nil
}
