package weather

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/arskydev/weatherman/internal/flags"
)

type Weather struct {
	Temperature     float64
	WeatherType     string
	Sunrise, Sunset time.Time
	Country, City   string
}

func (w *Weather) MarshalJSON() ([]byte, error) {
	jsonWeather := struct {
		Location struct {
			Country string `json:"country,omitempty"`
			Flag    string `json:"flag,omitempty"`
			City    string `json:"city,omitempty"`
		} `json:"location"`
		Weather struct {
			Temperature       float64 `json:"temperature,omitempty"`
			Unit              string  `json:"unit,omitempty"`
			WeatherType       string  `json:"weatherType,omitempty"`
			WeatherTypeSymbol string  `json:"weathersymbol,omitempty"`
		} `json:"weather"`
		Suncycle struct {
			Sunrise string `json:"sunrise,omitempty"`
			Sunset  string `json:"sunset,omitempty"`
		} `json:"suncycle"`
	}{
		Location: struct {
			Country string `json:"country,omitempty"`
			Flag    string `json:"flag,omitempty"`
			City    string `json:"city,omitempty"`
		}{
			Country: w.Country,
			Flag:    flags.GetFlag(w.Country),
			City:    w.City,
		},
		Weather: struct {
			Temperature       float64 `json:"temperature,omitempty"`
			Unit              string  `json:"unit,omitempty"`
			WeatherType       string  `json:"weatherType,omitempty"`
			WeatherTypeSymbol string  `json:"weathersymbol,omitempty"`
		}{
			Temperature:       w.Temperature,
			Unit:              "℃",
			WeatherType:       w.WeatherType,
			WeatherTypeSymbol: GetWeatherTypeSymbol(w),
		},
		Suncycle: struct {
			Sunrise string `json:"sunrise,omitempty"`
			Sunset  string `json:"sunset,omitempty"`
		}{
			Sunrise: w.Sunrise.Format(time.RFC822Z),
			Sunset:  w.Sunset.Format(time.RFC822Z),
		},
	}
	return json.Marshal(jsonWeather)
}

func (w *Weather) String() string {
	weatherTypeSymbol := GetWeatherTypeSymbol(w)
	flag := flags.GetFlag(w.Country)

	formatedWeather := fmt.Sprintf(
		`Weather for now:
Country: %v %v
City: %v
Temperature: %v ℃
Skies: %v %v
Sunrise: %v
Sunset: %v`,
		w.Country, flag, w.City, w.Temperature, w.WeatherType, weatherTypeSymbol, w.Sunrise.Format(time.RFC822Z), w.Sunset.Format(time.RFC822Z))

	return formatedWeather
}
