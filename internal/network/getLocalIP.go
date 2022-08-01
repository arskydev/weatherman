package network

import (
	"errors"

	"github.com/arskydev/weatherman/internal/requester"
)

var (
	url = "https://api.ipgeolocation.io/getip"
)

func GetLocalIP() (ip string, err error) {
	jsonResp, err := requester.GetJsonResp(url)

	if err != nil {
		return "", err
	}

	ip, ok := jsonResp["ip"].(string)

	if !ok {
		return "", errors.New("cannot get local ip. no such field in response")
	}

	return ip, nil
}
