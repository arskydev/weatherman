package network

import (
	"net/http"
	"strings"
)

func GetRemoteIp(r *http.Request) (ip string, err error) {
	ip = strings.Split(r.RemoteAddr, ":")[0]
	// If API server running localy
	if ip == "[" || ip[:3] == "192" || ip[:3] == "172" || ip[:3] == "127" || ip[:] == "10" {
		ip, err = GetLocalIP()
		if err != nil {
			return "", err
		}
	}

	return ip, nil
}
