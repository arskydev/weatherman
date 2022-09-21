package urls

import "net/url"

func SetURLQuery(u *url.URL, query map[string]string) {
	q := u.Query()
	for k, v := range query {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
}
