package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arskydev/weatherman/internal/coordinates"
	"github.com/arskydev/weatherman/internal/network"
	"github.com/arskydev/weatherman/internal/weather"
)

func main() {
	var (
		ipGeoKey      = os.Getenv("IPGEO_API_KEY")
		weatherApiKey = os.Getenv("WEATHER_API_KEY")
	)

	ip, err := network.GetLocalIP()

	if err != nil {
		log.Fatal("Error in network.GetLocalIP():\n", err)
	}

	coordinator := coordinates.New(ipGeoKey)
	weatherer := weather.NewWeatherer(coordinator, weatherApiKey)

	w, err := weatherer.GetWeather(ip)
	if err != nil {
		log.Fatal("Error in weather.GetWeather():\n", err)
	}

	// Output result to the console
	fmt.Println(w)
}
