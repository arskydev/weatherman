package requester

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

func GetJsonResp(url string) (jsonResp map[string]interface{}, err error) {
	var (
		resp        *http.Response
		c           = make(chan struct{})
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	)
	defer cancel()

	go func() {
		resp, err = http.Get(url)
		c <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return nil, errors.New("external service request timeout error")
	case <-c:
		if err != nil {
			return nil, err
		}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &jsonResp)

	return jsonResp, nil
}
