package weather

import (
	"fmt"
	"net/url"

	"github.com/arskydev/weatherman/internal/coordinates"
	"github.com/arskydev/weatherman/internal/requester"
	"github.com/arskydev/weatherman/pkg/web/urls"
)

type Weatherer struct {
	coordinator   *coordinates.Coordinator
	weatherApiKey string
}

func NewWeatherer(c *coordinates.Coordinator, weatherApiKey string) *Weatherer {
	return &Weatherer{
		coordinator:   c,
		weatherApiKey: weatherApiKey,
	}
}

func (w *Weatherer) GetWeather(ip string) (weather *Weather, err error) {
	var (
		weatherApiUrl = &url.URL{
			Scheme:     "https",
			Host:       "api.openweathermap.org",
			Path:       "/data/2.5/weather",
			ForceQuery: false,
		}
	)
	c, err := w.coordinator.Get(ip)

	if err != nil {
		return nil, fmt.Errorf("error on GetCoordinates: %w", err)
	}

	weatherApiUrlQuery := map[string]string{
		"appid": w.weatherApiKey,
		"lat":   c.Latitude,
		"lon":   c.Longitude,
		"units": "metric",
	}
	urls.SetURLQuery(weatherApiUrl, weatherApiUrlQuery)
	jsonResp := weatherRaw{}
	if err := requester.GetJson(weatherApiUrl.String(), &jsonResp); err != nil {
		return nil, err
	}

	weather = jsonResp.generateWeather()

	return weather, nil
}
