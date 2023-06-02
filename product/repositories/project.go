package repositories

import "photovoltaic-system-services/db"

func CheckExistProject(projectId int, userId int) int64 {
	var count int64
	db.Database.Table("projects").Select("count(id)").Where("id = ? and user_id = ?", projectId, userId).Count(&count)
	return count
}
