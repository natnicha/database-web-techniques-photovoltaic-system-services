package repositories

import (
	"photovoltaic-system-services/db"
	"time"
)

type Users struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  []byte    `json:"password"`
	IsActive  bool      `json:"is_active"`
	UpdateAt  time.Time `json:"update_at"`
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

func DeleteUserById(id string) (err error) {
	result := db.Database.Where("id = ?", id).Delete(Users{})
	if result.Error != nil {
		return err
	}
	return nil
}
