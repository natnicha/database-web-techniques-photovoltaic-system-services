package repositories

import (
	"fmt"
	"photovoltaic-system-services/db"
	"time"
)

type WeatherInfo struct {
	Count        int64     `json:"cnt"`
	StartWeather time.Time `json:"start_weather"`
	EndWeather   time.Time `json:"end_weather"`
	Latitude     string    `json:"latitude"`
	Longtitude   string    `json:"longtitude"`
}

func GetWeatherInfo(projectId int) WeatherInfo {
	query := `SELECT count(*) as cnt, Timezone('Europe/Berlin',(cast(TO_CHAR(pj.start_at - INTERVAL '30 day', 'YYYY-MM-DD HH24:00:00') as timestamp))) as start_weather, Timezone('Europe/Berlin',cast(cast(pj.start_at as date) as timestamp)) as end_weather, w.latitude, w.longitude
		FROM projects pj 
		LEFT JOIN products p ON p.project_id = pj.id
		LEFT JOIN weather w ON w.latitude = p.geolocation[0] AND w.longitude = p.geolocation[1]
		WHERE p.id = ` + fmt.Sprint(projectId) + `
		AND (w.datetime >= (cast(TO_CHAR(pj.start_at - INTERVAL '30 day', 'YYYY-MM-DD HH24:00:00') as timestamp)) ) AND (w.datetime < (cast((pj.start_at) as date)))
		GROUP BY pj.start_at, w.latitude, w.longitude
	`
	var result WeatherInfo
	row := db.Database.Raw(query).Row()
	row.Scan(&result.Count, &result.StartWeather, &result.EndWeather, &result.Latitude, &result.Longtitude)
	return result
}
