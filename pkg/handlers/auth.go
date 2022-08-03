package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/arskydev/weatherman/pkg/users"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var user users.User
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		msg := "Invalid request body"
		h.sendErrorResponse(msg, http.StatusBadRequest, w, err)
		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		msg := "Invalid username, email and password passed"
		h.sendErrorResponse(msg, http.StatusBadRequest, w, err)
		return
	}

	id, err := h.service.Authorization.CreateUser(user)

	if err != nil {
		msg := "Error while creating user: " + err.Error()
		h.sendErrorResponse(msg, http.StatusInternalServerError, w, err)
		return
	}

	resp := map[string]string{"id": strconv.Itoa(id)}
	h.sendSimpleResponse(resp, http.StatusCreated, w)
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var user users.User
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		msg := "Invalid request body"
		h.sendErrorResponse(msg, http.StatusBadRequest, w, err)
		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		msg := "Invalid username and password"
		h.sendErrorResponse(msg, http.StatusBadRequest, w, err)
		return
	}

	token, err := h.service.Authorization.GenerateToken(user.Username, user.Password)

	if err != nil {
		msg := "Error while creating user: " + err.Error()
		h.sendErrorResponse(msg, http.StatusInternalServerError, w, err)
		return
	}

	resp := map[string]string{"token": token}
	h.sendSimpleResponse(resp, http.StatusOK, w)
}
