package handlers

import (
	"errors"
	"net/http"
)

const (
	AUTH_HEADER_NAME = "Authorization"
)

func (h *Handler) ValidateJWT(next func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Connection", "close")
		defer r.Body.Close()
		if r.Header[AUTH_HEADER_NAME] == nil {
			msg := "empty token"
			h.sendErrorResponse(msg, http.StatusUnauthorized, w, errors.New(msg))
			return
		}

		token, err := h.service.ValidateToken(r.Header[AUTH_HEADER_NAME][0])

		if err != nil {
			msg := "Access denied"
			h.sendErrorResponse(msg, http.StatusForbidden, w, err)
			return
		}

		if token.Valid {
			next(w, r)
		}

	})
}
