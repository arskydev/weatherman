package responder

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonResponse(resp map[string]string, statusCode int, w http.ResponseWriter) {
	js, err := json.Marshal(resp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unexpected error while marshalling json..."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func SendJSONResponse(resp map[string]string, statusCode int, w http.ResponseWriter) {
	jsonResponse(resp, statusCode, w)
}

func SendErrorResponse(msg string, statusCode int, w http.ResponseWriter, err error) {
	resp := map[string]string{"msg": msg}
	jsonResponse(resp, statusCode, w)
	log.Println("Error while reading request body", err)
}
