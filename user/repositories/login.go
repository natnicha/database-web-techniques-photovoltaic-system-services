package repositories

import (
	"photovoltaic-system-services/db"
)

func CreateLogin(id int) (err error) {
	result := db.Database.Exec("INSERT INTO Login (user_id) Values (?)", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
