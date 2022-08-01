package url

import "fmt"

func BuildURL(baseURL string, args ...interface{}) string {
	return fmt.Sprintf(baseURL, args...)
}
