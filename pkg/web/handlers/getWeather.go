package handlers

import (
	"errors"
	"net/http"

	"github.com/arskydev/weatherman/internal/formater"
	"github.com/arskydev/weatherman/internal/network"
	"github.com/arskydev/weatherman/internal/weather"
	"github.com/arskydev/weatherman/pkg/web/internal/responder"
)

func (h *Handler) getWeather(w http.ResponseWriter, r *http.Request) {
	ip, err := network.GetRemoteIp(r)

	if err != nil {
		msg := "error while getting user IP"
		statusCode := http.StatusInternalServerError
		responder.ErrorSampleTextResponse(msg, statusCode, w, errors.New(msg))
		return
	}

	_weather, err := weather.GetWeather(ip)

	if err != nil {
		msg := "error while getting weather"
		statusCode := http.StatusInternalServerError
		responder.ErrorSampleTextResponse(msg, statusCode, w, errors.New(msg))
		return
	}

	j, err := formater.FormatWeatherJson(_weather)

	if err != nil {
		msg := "error while formating JSON response"
		statusCode := http.StatusInternalServerError
		responder.ErrorSampleTextResponse(msg, statusCode, w, errors.New(msg))
		return
	}

	w.Write(j)
}
