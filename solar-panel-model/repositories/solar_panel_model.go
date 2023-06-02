package repositories

import (
	"photovoltaic-system-services/db"
	"strings"
)

type SolarPanelModel struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Efficiency  float32 `json:"efficiency"`
}

type ListRequest struct {
	Filter  string `form:"filter"`
	Limit   int    `form:"limit" binding:"required,min=1"`
	Offset  int    `form:"offset" binding:"required"`
	OrderBy string `form:"order_by"`
	SortBy  string `form:"sort_by"`
}

func GetSolarPanelModel(query ListRequest) (*[]SolarPanelModel, error) {
	tx := db.Database.Model(&SolarPanelModel{})
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

	projects := []SolarPanelModel{}
	result := tx.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return &projects, nil
}
