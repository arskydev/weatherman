package middleware

import (
	"errors"
	"net/http"

	"github.com/arskydev/weatherman/pkg/web/internal/responder"
)

const (
	AUTH_HEADER_NAME = "Authorization"
)

func (m *Middleware) ValidateJWT(next func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Connection", "close")
		defer r.Body.Close()
		if r.Header[AUTH_HEADER_NAME] == nil {
			msg := "empty token"
			responder.SendErrorResponse(msg, w, errors.New(msg))
			return
		}

		token, err := m.auth.ValidateToken(r.Header[AUTH_HEADER_NAME][0])

		if err != nil {
			msg := "Access denied"
			responder.SendErrorResponse(msg, w, err)
			return
		}

		if token.Valid {
			next(w, r)
		}

	})
}
