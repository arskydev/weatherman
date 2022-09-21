package weather

import (
	"time"
)

type weatherRaw struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	WeatherList weatherList `json:"weather"`
	City        string      `json:"name"`
}

func (w *weatherRaw) generateWeather() *Weather {
	temperature := w.Main.Temp
	sunrise := time.Unix(w.Sys.Sunrise, 0)
	sunset := time.Unix(w.Sys.Sunset, 0)
	city := w.City
	country := w.Sys.Country
	weatherType := w.WeatherList[0].Main

	return &Weather{
		Temperature: temperature,
		WeatherType: weatherType,
		Sunrise:     sunrise,
		Sunset:      sunset,
		Country:     country,
		City:        city,
	}
}
