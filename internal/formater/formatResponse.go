package formater

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/arskydev/weatherman/internal/flags"
	"github.com/arskydev/weatherman/internal/weather"
)

func FormatWeatherString(w *weather.Weather) (formatedWeather string) {
	weatherTypeSymbol := weather.GetWeatherTypeSymbol(w)
	flag := flags.GetFlag(w.Country)

	formatedWeather = fmt.Sprintf(
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

func FormatWeatherJson(w *weather.Weather) ([]byte, error) {
	//tbh, gophers are not lazy beings. And we do like this:
	//jsonWeather := struct {
	//	Location struct {
	//		Country string `json:"country,omitempty"`
	//		Flag    string `json:"flag,omitempty"`
	//		City    string `json:"city,omitempty"`
	//	} `json:"location"`
	//	Weather struct {
	//		Temperature float64 `json:"temperature,omitempty"`
	//		WeatherType string  `json:"weatherType,omitempty"`
	//		//}
	//	} `json:"weather"`
	//
	//}{}
	// etc

	jsonWeather := struct {
		Location map[string]interface{} `json:"location"`
		// country     string
		// flag        string
		// city        string
		Weather map[string]interface{} `json:"weather"`
		// temperature float64
		// weatherType string
		// weatherSymbol string
		Suncycle map[string]interface{} `json:"suncycle"`
		// sunrise     string
		// sunset      string
	}{
		Location: map[string]interface{}{
			"country": w.Country,
			"flag":    flags.GetFlag(w.Country),
			"city":    w.City,
		},
		Weather: map[string]interface{}{
			"temperature":   w.Temperature,
			"unit":          "℃",
			"weathertype":   w.WeatherType,
			"weathersymbol": weather.GetWeatherTypeSymbol(w),
		},
		Suncycle: map[string]interface{}{
			"sunrise": w.Sunrise.Format(time.RFC822Z),
			"sunset":  w.Sunset.Format(time.RFC822Z),
		},
	}
	return json.Marshal(jsonWeather)
}
