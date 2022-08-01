package weather

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/arskydev/weatherman/internal/coordinates"
	"github.com/arskydev/weatherman/internal/requester"
	"github.com/arskydev/weatherman/internal/url"
)

const (
	WEATHERAPI_URL_BASE = "https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v&units=metric"
)

type Weather struct {
	Temperature     float64
	WeatherType     string
	Sunrise, Sunset time.Time
	Country, City   string
}

func GetWeather(ip string) (weather *Weather, err error) {
	c, err := coordinates.GetCoordinates(ip)

	if err != nil {
		return nil, fmt.Errorf("error on GetCoordinates: %w", err)
	}

	weatherApiKey := os.Getenv("WEATHER_API_KEY")

	if weatherApiKey == "" {
		return nil, errors.New("no WEATHER_API_KEY passed")
	}

	weatherArgs := []interface{}{c.Latitude, c.Longitude, weatherApiKey}
	url := url.BuildURL(WEATHERAPI_URL_BASE, weatherArgs...)
	jsonResp, err := requester.GetJsonResp(url)

	if err != nil {
		return weather, err
	}

	weather = generateWeather(jsonResp)

	return weather, nil
}

func jsonRespParser(jsonResp map[string]interface{}, key string) map[string]interface{} {
	switch key {

	case "main", "sys":
		return jsonResp[key].(map[string]interface{})

	case "weather":
		w := jsonResp[key].([]interface{})[0]
		return w.(map[string]interface{})

	default:
		return nil
	}
}

func generateWeather(jsonResp map[string]interface{}) *Weather {
	mainJson := jsonRespParser(jsonResp, "main")
	sysJson := jsonRespParser(jsonResp, "sys")
	weatherJson := jsonRespParser(jsonResp, "weather")

	temperature := mainJson["temp"].(float64)
	sunrise_raw := int64(sysJson["sunrise"].(float64))
	sunrise := time.Unix(sunrise_raw, 0)
	sunset_raw := int64(sysJson["sunset"].(float64))
	sunset := time.Unix(sunset_raw, 0)
	city := jsonResp["name"].(string)
	country := sysJson["country"].(string)
	weatherType := weatherJson["main"].(string)

	return &Weather{Temperature: temperature, WeatherType: weatherType, Sunrise: sunrise, Sunset: sunset, Country: country, City: city}
}
