package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) sendSimpleResponse(resp map[string]string, statusCode int, w http.ResponseWriter) {
	j, err := json.Marshal(resp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unexpected error while marshalling json..."))
		return
	}

	w.WriteHeader(statusCode)
	w.Write(j)
}

func (h *Handler) sendErrorResponse(msg string, statusCode int, w http.ResponseWriter, err error) {
	resp := map[string]string{"msg": msg}
	w.WriteHeader(statusCode)
	h.sendSimpleResponse(resp, statusCode, w)
	log.Println("Error while reading request body", err)
}
