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
	query := `WITH project_info as (
		SELECT Timezone('Europe/Berlin',(cast(TO_CHAR(pj.start_at - INTERVAL '30 day', 'YYYY-MM-DD HH24:00:00') as timestamp))) as start_weather, Timezone('Europe/Berlin',cast(cast(pj.start_at as date) as timestamp)) as end_weather, p.geolocation[0] as latitude, p.geolocation[1] as longitude
		FROM projects pj 
		LEFT JOIN products p ON p.project_id = pj.id
		WHERE p.id = ` + fmt.Sprint(projectId) + `
	)
	
	SELECT weather.cnt, pi.* 
	FROM project_info pi
	LEFT JOIN (
		SELECT count(*) as cnt
		FROM project_info pi
		LEFT JOIN weather w  ON w.latitude = pi.latitude AND w.longitude = pi.longitude
		WHERE (w.datetime >= pi.start_weather ) AND (w.datetime < pi.end_weather)
	)AS weather ON 1=1
	`
	var result WeatherInfo
	row := db.Database.Raw(query).Row()
	row.Scan(&result.Count, &result.StartWeather, &result.EndWeather, &result.Latitude, &result.Longtitude)
	return result
}
