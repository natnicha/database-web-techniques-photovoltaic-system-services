package repositories

import (
	"photovoltaic-system-services/db"
)

type Users struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"password"`
	IsActive  bool   `json:"is_active"`
}

func GetUserById(id string) (user *Users, err error) {
	result := db.Database.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user Users) (*Users, error) {
	result := db.Database.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func UpdateUser(id string, user Users) error {
	tx := db.Database.Where("id = ?", id).Updates(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
