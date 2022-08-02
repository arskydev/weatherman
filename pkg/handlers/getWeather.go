package handlers

import (
	"net/http"

	"github.com/arskydev/weatherman/internal/formater"
	"github.com/arskydev/weatherman/internal/network"
	"github.com/arskydev/weatherman/internal/weather"
)

func (h *Handler) getWeather(w http.ResponseWriter, r *http.Request) {
	ip, err := network.GetRemoteIp(r)

	if err != nil {
		msg := "Error while getting user IP"
		h.sendErrorResponse(msg, http.StatusInternalServerError, w, err)
		return
	}

	_weather, err := weather.GetWeather(ip)

	if err != nil {
		msg := "Error while getting weather"
		h.sendErrorResponse(msg, http.StatusInternalServerError, w, err)
		return
	}

	j, err := formater.FormatWeatherJson(_weather)

	if err != nil {
		msg := "Error while formating JSON response"
		h.sendErrorResponse(msg, http.StatusInternalServerError, w, err)
		return
	}

	w.Write(j)
}
