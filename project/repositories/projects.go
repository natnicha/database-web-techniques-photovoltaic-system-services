package repositories

import (
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
}

func CreateProject(project Projects) (*Projects, error) {
	result := db.Database.Create(&project)
	if result.Error != nil {
		return nil, result.Error
	}
	return &project, nil
}
