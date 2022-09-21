package requester

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func GetJson(url string, object any) error {
	var (
		err  error
		resp *http.Response
	)
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err = client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	json.Unmarshal(body, &object)

	return nil
}
