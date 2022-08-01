package handlers

import (
	"fmt"
	"net/http"

	"github.com/arskydev/weatherman/internal/formater"
	"github.com/arskydev/weatherman/internal/network"
	"github.com/arskydev/weatherman/internal/weather"
)

func (h *Handler) getWeather(w http.ResponseWriter, r *http.Request) {
	ip, err := network.GetRemoteIp(r)

	if err != nil {
		msg := fmt.Sprint("Error on GetRemoteIp:\n", err)
		w.Write([]byte(msg))
		return
	}

	_weather, err := weather.GetWeather(ip)

	if err != nil {
		msg := fmt.Sprint("Error on GetWeather:\n", err)
		w.Write([]byte(msg))
		return
	}

	j, err := formater.FormatWeatherJson(_weather)

	if err != nil {
		msg := fmt.Sprint("Error on FormatWeatherJson:\n", err)
		w.Write([]byte(msg))
		return
	}

	w.Write(j)
}
