package network

import (
	"errors"
	"fmt"

	"github.com/arskydev/weatherman/internal/requester"
)

func GetLocalIP() (ip string, err error) {
	var (
		url = "https://api.ipgeolocation.io/getip"
	)
	jsonResp := &struct {
		Ip string `json:"ip"`
	}{}
	err = requester.GetJson(url, jsonResp)
	fmt.Println()

	if err != nil {
		return "", err
	}

	if jsonResp.Ip == "" {
		return "", errors.New("cannot get local ip. no such field in response")
	}

	return jsonResp.Ip, nil
}
