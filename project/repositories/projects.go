package repositories

import (
	"errors"
	"photovoltaic-system-services/db"
	"time"
)

type Projects struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"start_at"`
	IsPrinted   bool      `json:"is_printed"`
	UpdateAt    time.Time `json:"update_at"`
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

func UpdateProject(projectId int, project Projects) (*Projects, error) {
	project.Id = projectId
	tx := db.Database.Model(&project).Where("id = ? and user_id = ?", projectId, project.UserId)
	var count int64
	tx.Count(&count)
	if count == 0 {
		return nil, errors.New("No project ID with specified user ID ")
	}
	tx.Save(project)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &project, nil
}

func DeleteProjectById(id int) (err error) {
	result := db.Database.Where("id = ?", id).Delete(Projects{})
	if result.Error != nil {
		return err
	}
	return nil
}
