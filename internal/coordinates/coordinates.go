package coordinates

import (
	"github.com/arskydev/weatherman/internal/requester"
	"github.com/arskydev/weatherman/internal/urlBuilder"
)

const (
	ipGeoUrlBase = "https://api.ipgeolocation.io/ipgeo?apiKey=%v&ip=%v"
)

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Coordinator struct {
	ipGEOKey string
}

func New(ipGeoKey string) *Coordinator {
	return &Coordinator{
		ipGEOKey: ipGeoKey,
	}
}

// this is just an example how can we avoid using os.Getenv("IPGEO_API_KEY") in this package
func (c *Coordinator) Get(ip string) (*Coordinates, error) {

	ipGEOurl := urlBuilder.BuildURL(ipGeoUrlBase, c.ipGEOKey, ip)
	coords := &Coordinates{}

	if err := requester.GetJson(ipGEOurl, coords); err != nil {
		return nil, err
	}

	return coords, nil
}
