package coordinates

import (
	"errors"
	"github.com/arskydev/weatherman/internal/requester"
	"github.com/arskydev/weatherman/internal/url"
	"os"
)

const (
	IP_GEO_URL_BASE = "https://api.ipgeolocation.io/ipgeo?apiKey=%v&ip=%v"
)

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Coordinator struct {
	ipGEOKey string
}

func New(ipGeoKey string) Coordinator {
	return Coordinator{
		ipGEOKey: ipGeoKey,
	}
}

//this is just an example how can we avoid using os.Getenv("IPGEO_API_KEY") in this package
func (c *Coordinator) Get(ip string) (*Coordinates, error) {
	return nil, nil
}

func GetCoordinates(ip string) (*Coordinates, error) {
	ipgeoApiKey := os.Getenv("IPGEO_API_KEY") // This should be moved out of here

	if ipgeoApiKey == "" {
		return nil, errors.New("no IPGEO_API_KEY passed")
	}

	//coordArgs := []interface{}{ipgeoApiKey, ip}
	// this is a bit more readable. It's OK to unpack if we get the slice from somewhere
	//url name collides with package name. We want to avoid it
	ipGEOurl := url.BuildURL(IP_GEO_URL_BASE, ipgeoApiKey, ip)
	//Sorry, I just had to rework it a bit. Linus said "Words are cheap, show me the code". I'm showing.
	c := &Coordinates{}
	if err := requester.GetJson(ipGEOurl, c); err != nil {
		return nil, err
	}
	return c, nil
}
