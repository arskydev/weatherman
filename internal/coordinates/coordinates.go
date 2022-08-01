package coordinates

import (
	"errors"
	"os"
	"strconv"

	"github.com/arskydev/weatherman/internal/requester"
	"github.com/arskydev/weatherman/internal/url"
)

const (
	IP_GEO_URL_BASE = "https://api.ipgeolocation.io/ipgeo?apiKey=%v&ip=%v"
)

type Coordinates struct {
	Latitude, Longitude float64
}

func GetCoordinates(ip string) (*Coordinates, error) {
	ipgeoApiKey := os.Getenv("IPGEO_API_KEY")

	if ipgeoApiKey == "" {
		return nil, errors.New("no IPGEO_API_KEY passed")
	}

	coordArgs := []interface{}{ipgeoApiKey, ip}
	url := url.BuildURL(IP_GEO_URL_BASE, coordArgs...)
	jsonResp, err := requester.GetJsonResp(url)

	if err != nil {
		return nil, err
	}

	latt, ok := jsonResp["latitude"].(string)

	if !ok {
		return nil, errors.New("no lattitude in response")
	} else {
		latt = latt[:4]
	}

	longit, ok := jsonResp["longitude"].(string)

	if !ok {
		return nil, errors.New("no longitude in response")
	} else {
		longit = longit[:4]
	}

	latitude, err := strconv.ParseFloat(latt, 64)

	if err != nil {
		return nil, err
	}

	longitude, err := strconv.ParseFloat(longit, 64)

	if err != nil {
		return nil, err
	}

	return &Coordinates{Latitude: latitude, Longitude: longitude}, nil
}
