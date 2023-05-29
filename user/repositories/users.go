package repositories

import (
	"photovoltaic-system-services/db"
)

type User interface {
	GetUserById(id int) (user *users, err error)
}

type users struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsActive  bool   `json:"is_active"`
}

func GetUserById(id string) (user []users, err error) {
	result := db.Database.Where("id = ?", id).Find(&user)
	_, err = result.Rows()
	if err != nil {
		return nil, err
	}
	return user, nil
}
