package repositories

import (
	"photovoltaic-system-services/db"
)

type UserLog struct {
	Type      string
	UserId    int
	Host      string
	UserAgent string
}

func InsertLoginUserLog(userLog UserLog) (err error) {
	tx := db.Database.Create(&userLog)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
