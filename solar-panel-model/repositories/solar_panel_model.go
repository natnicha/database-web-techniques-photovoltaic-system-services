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
		if strings.Contains(query.Filter, "&") {
			filters := strings.Split(query.Filter, "&")
			where := ""
			for _, filter := range filters {
				if where != "" {
					where += " and "
				}
				f := strings.Split(filter, ":")
				if f[0] == "name" {
					where += "cast( " + f[0] + " as varchar)  like '%" + f[1] + "%'"
				} else {
					where += f[0] + "=" + f[1]
				}
			}
			tx.Where(where)
		} else {
			filters := strings.Split(query.Filter, ":")
			if filters[0] == "name" {
				tx.Where("cast( " + filters[0] + " as varchar)  like '%" + filters[1] + "%'")
			} else {
				tx.Where(filters[0] + "=" + filters[1])
			}
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

	projects := []SolarPanelModel{}
	result := tx.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return &projects, nil
}
