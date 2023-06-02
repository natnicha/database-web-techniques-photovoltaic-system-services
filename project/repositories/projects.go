package repositories

import (
	"photovoltaic-system-services/db"
	"strings"
	"time"
)

type Projects struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"start_at"`
	IsPrinted   bool      `json:"is_printed"`
}

type ListRequest struct {
	Filter  string `form:"filter"`
	Limit   int    `form:"limit" binding:"required,min=1"`
	Offset  int    `form:"offset" binding:"required"`
	OrderBy string `form:"order_by"`
	SortBy  string `form:"sort_by"`
}

func CreateProject(project Projects) (*Projects, error) {
	result := db.Database.Create(&project)
	if result.Error != nil {
		return nil, result.Error
	}
	return &project, nil
}

func CheckExistProject(projectId int, project Projects) int64 {
	var count int64
	_ = db.Database.Model(&project).Where("id = ? and user_id = ?", projectId, project.UserId).Count(&count)
	return count
}

func DeleteProjectById(id int) (err error) {
	result := db.Database.Where("id = ?", id).Delete(Projects{})
	if result.Error != nil {
		return err
	}
	return nil
}

func GetProject(query ListRequest) (*[]Projects, error) {
	tx := db.Database.Model(&Projects{})
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
	projects := []Projects{}
	result := tx.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return &projects, nil
}

func UpdateProject(projectId int, project Projects) (*Projects, error) {
	project.Id = projectId
	tx := db.Database.Model(&project).Where("id = ? and user_id = ?", projectId, project.UserId).UpdateColumn("update_at", "now()").Save(project)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &project, nil
}
