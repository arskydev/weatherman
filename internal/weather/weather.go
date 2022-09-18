package weather

import (
	"fmt"
	"time"

	"github.com/arskydev/weatherman/internal/coordinates"
	"github.com/arskydev/weatherman/internal/requester"
	"github.com/arskydev/weatherman/internal/urlBuilder"
)

const (
	weatherApiUrlBase = "https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v&units=metric"
)

type Weather struct {
	Temperature     float64
	WeatherType     string
	Sunrise, Sunset time.Time
	Country, City   string
}

type Weatherer struct {
	coordinator   *coordinates.Coordinator
	weatherApiKey string
}

func New(c *coordinates.Coordinator, weatherApiKey string) *Weatherer {
	return &Weatherer{
		coordinator:   c,
		weatherApiKey: weatherApiKey,
	}
}

func (w *Weatherer) GetWeather(ip string) (weather *Weather, err error) {
	c, err := w.coordinator.Get(ip)

	if err != nil {
		return nil, fmt.Errorf("error on GetCoordinates: %w", err)
	}

	weatherArgs := []interface{}{c.Latitude, c.Longitude, w.weatherApiKey}
	url := urlBuilder.BuildURL(weatherApiUrlBase, weatherArgs...)
	jsonResp, err := requester.GetJsonResp(url)

	if err != nil {
		return nil, err
	}

	weather = generateWeather(jsonResp)

	return weather, nil
}

// These two methods should get the logic from internal/formater/formatResponse.go
func (w *Weather) MarshalJSON() ([]byte, error) {
	return []byte("{\"here you can put some stuff\"})"), nil
}

func (w *Weather) String() string {
	return "now you can print me!"
}

// please avoid using map[string]interface{} when unmarshalling json.
// @see internal/requester/requester.go
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
