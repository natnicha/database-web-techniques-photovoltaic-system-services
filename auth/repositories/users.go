package repositories

import (
	"photovoltaic-system-services/db"
)

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}

func GetUserById(id int) (user *Users, err error) {
	result := db.Database.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByEmail(email string) (user *Users, err error) {
	result := db.Database.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, err
	}
	return user, nil
}
