package coordinates

import (
	"net/url"

	"github.com/arskydev/weatherman/internal/requester"
	"github.com/arskydev/weatherman/pkg/web/urls"
)

type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Coordinator struct {
	ipGEOKey string
}

func New(ipGeoKey string) *Coordinator {
	return &Coordinator{
		ipGEOKey: ipGeoKey,
	}
}

func (c *Coordinator) Get(ip string) (*Coordinates, error) {
	var (
		ipGeoUrl = &url.URL{
			Scheme:     "https",
			Host:       "api.ipgeolocation.io",
			Path:       "/ipgeo",
			ForceQuery: false,
		}
	)
	ipGeoUrlQuery := map[string]string{
		"apiKey": c.ipGEOKey,
		"ip":     ip,
	}
	urls.SetURLQuery(ipGeoUrl, ipGeoUrlQuery)
	coords := Coordinates{}

	if err := requester.GetJson(ipGeoUrl.String(), &coords); err != nil {
		return nil, err
	}

	coords.Latitude = coords.Latitude[:4]
	coords.Longitude = coords.Longitude[:4]
	return &coords, nil
}
