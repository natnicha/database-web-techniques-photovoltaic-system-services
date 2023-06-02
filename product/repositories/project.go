package repositories

import "photovoltaic-system-services/db"

func CheckExistProject(projectId int, userId int) int64 {
	var count int64
	db.Database.Table("projects").Select("count(id)").Where("id = ? and user_id = ?", projectId, userId).Count(&count)
	return count
}

func CheckExistProjectByProductIdAndUserId(productId int, userId int) int64 {
	var count int64
	statement := "SELECT count(pd.id) FROM products pd LEFT JOIN projects pj ON pd.project_id = pj.id WHERE pd.id = ? AND pj.user_id = ?"
	db.Database.Raw(statement, productId, userId).Count(&count)
	return count
}
