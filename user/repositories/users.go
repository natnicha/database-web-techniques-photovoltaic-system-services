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
}

func CreateUser(user Users) (*Users, error) {
	result := db.Database.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func DeleteUserById(id string) (err error) {
	result := db.Database.Where("id = ?", id).Delete(Users{})
	if result.Error != nil {
		return err
	}
	return nil
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

func UpdateUser(id int, user Users) (*Users, error) {
	user.Id = id
	tx := db.Database.Model(&user).Where("id = ?", id).Save(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = db.Database.Model(&user).Where("id = ?", id).UpdateColumn("update_at", "now()")
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
