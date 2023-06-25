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
	Longitude    string    `json:"longtitude"`
}

func GetWeatherInfo(projectId int) WeatherInfo {
	query := `WITH report_range as (
		Select (CAST((pj.start_at at time zone 'utc' + INTERVAL '30 day') as date) <= CAST(now() as date)) as is_over_30_days
		,pj.start_at at time zone 'utc' as start_at, p.geolocation[0] as latitude, p.geolocation[1] as longitude
		FROM projects pj 
		LEFT JOIN products p ON p.project_id = pj.id
		WHERE p.id = ` + fmt.Sprint(projectId) + `
	)
	, project_info as (
		SELECT 
			Case when is_over_30_days
				then cast(TO_CHAR(r.start_at , 'YYYY-MM-DD HH24:00:00') as timestamp)
				else cast(cast(now() at time zone 'utc' as date) - INTERVAL '30 day' as timestamp) 
			end as start_weather,
			Case when is_over_30_days
				then cast(cast((r.start_at + INTERVAL '30 day') as date) as timestamp) - INTERVAL '1 second'
				else cast(cast(now() at time zone 'utc' as date) as timestamp) - INTERVAL '1 second'
			end as end_weather
			, r.latitude, r.longitude
		FROM report_range r
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
	row.Scan(&result.Count, &result.StartWeather, &result.EndWeather, &result.Latitude, &result.Longitude)
	return result
}
