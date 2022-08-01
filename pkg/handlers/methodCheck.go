package handlers

import (
	"errors"
	"net/http"
)

func checkMethod(r *http.Request, allowedMethods []string) error {
	isInList := false

	for _, method := range allowedMethods {
		if r.Method == method {
			isInList = true
		}
	}

	if !isInList {
		return errors.New("405 Method not allowed")
	}

	return nil
}
