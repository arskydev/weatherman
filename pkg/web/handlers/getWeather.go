package handlers

import (
	"net/http"

	"github.com/arskydev/weatherman/internal/formater"
	"github.com/arskydev/weatherman/internal/network"
	"github.com/arskydev/weatherman/internal/weather"
	"github.com/arskydev/weatherman/pkg/web/internal/responder"
)

func (h *Handler) getWeather(w http.ResponseWriter, r *http.Request) {
	ip, err := network.GetRemoteIp(r)

	if err != nil {
		msg := "Error while getting user IP"
		responder.SendErrorResponse(msg, w, err)
		return
	}

	_weather, err := weather.GetWeather(ip)

	if err != nil {
		msg := "Error while getting weather"
		responder.SendErrorResponse(msg, w, err)
		return
	}

	j, err := formater.FormatWeatherJson(_weather)

	if err != nil {
		msg := "Error while formating JSON response"
		responder.SendErrorResponse(msg, w, err)
		return
	}

	w.Write(j)
}
