package responder

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonResp(resp map[string]string, statusCode int, w http.ResponseWriter) {
	j, err := json.Marshal(resp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unexpected error while marshalling json..."))
		return
	}

	w.WriteHeader(statusCode)
	w.Write(j)
}

func JSONResp(resp map[string]string, statusCode int, w http.ResponseWriter) {
	jsonResp(resp, statusCode, w)
}

func SendErrorResponse(msg string, statusCode int, w http.ResponseWriter, err error) {
	resp := map[string]string{"msg": msg}
	w.WriteHeader(statusCode)
	jsonResp(resp, statusCode, w)
	log.Println("Error while reading request body", err)
}
