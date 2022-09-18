package requester

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

//map[string]interface{} - the first thing we want to try with go when read JSON, but not the best(example will follow)
func GetJsonResp(url string) (jsonResp map[string]interface{}, err error) {
	var (
		resp        *http.Response
		c           = make(chan struct{})
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	)
	defer cancel()

	go func() {
		// usage of errgroup is advised to catch errors from goroutines.
		// Still, in this case we can do a bit simple approach (example below)
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
