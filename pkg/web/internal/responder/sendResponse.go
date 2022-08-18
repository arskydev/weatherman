package responder

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonResponse(resp map[string]string, w http.ResponseWriter) {
	js, err := json.Marshal(resp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unexpected error while marshalling json..."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func sampleTextResponse(msg string, statusCode int, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	w.Write([]byte(msg))
}

func ErrorSampleTextResponse(msg string, statusCode int, w http.ResponseWriter, err error) {
	sampleTextResponse(msg, statusCode, w)
	log.Println(err)
}

func SendJSONResponse(resp map[string]string, w http.ResponseWriter) {
	jsonResponse(resp, w)
}

func SendErrorJSONResponse(msg string, w http.ResponseWriter, err error) {
	resp := map[string]string{"msg": msg}
	jsonResponse(resp, w)
	log.Println("Error while reading request body", err)
}
