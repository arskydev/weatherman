package main

import (
	"fmt"
	"log"

	"github.com/arskydev/weatherman/internal/formater"
	"github.com/arskydev/weatherman/internal/network"
	"github.com/arskydev/weatherman/internal/weather"
)

func main() {
	ip, err := network.GetLocalIP()

	if err != nil {
		log.Fatal("Error in network.GetLocalIP():\n", err)
	}

	w, err := weather.GetWeather(ip)

	if err != nil {
		log.Fatal("Error in weather.GetWeather():\n", err)
	}

	result := formater.FormatWeatherString(w)

	// Output result to the console
	fmt.Println(result)
}
