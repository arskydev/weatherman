package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/arskydev/weatherman/pkg/users"
	"github.com/arskydev/weatherman/pkg/web/internal/responder"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var user users.User
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		msg := "Invalid request body"
		responder.SendErrorResponse(msg, w, err)
		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		msg := "Invalid username, email and password passed"
		responder.SendErrorResponse(msg, w, err)
		return
	}

	id, err := h.service.Authorization.CreateUser(user)

	if err != nil {
		msg := "Error while creating user: " + err.Error()
		responder.SendErrorResponse(msg, w, err)
		return
	}

	resp := map[string]string{"id": strconv.Itoa(id)}
	responder.SendJSONResponse(resp, w)
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var user users.User
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		msg := "Invalid request body"
		responder.SendErrorResponse(msg, w, err)
		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		msg := "Invalid username and password"
		responder.SendErrorResponse(msg, w, err)
		return
	}

	token, err := h.service.Authorization.GenerateToken(user.Username, user.Password)

	if err != nil {
		msg := "Error while creating user: " + err.Error()
		responder.SendErrorResponse(msg, w, err)
		return
	}

	resp := map[string]string{"token": token}
	responder.SendJSONResponse(resp, w)
}
