package urlBuilder

import (
	"fmt"
)

func BuildURL(baseURL string, args ...interface{}) string {
	//This might seem to be an overkill. But this is the DAO. It is safe, reliable and please don't use fmt to print URLs
	//urlbuilder := url.URL{
	//	Scheme:      "",
	//	Opaque:      "",
	//	User:        nil,
	//	Host:        "",
	//	Path:        "",
	//	RawPath:     "",
	//	ForceQuery:  false,
	//	RawQuery:    "",
	//	Fragment:    "",
	//	RawFragment: "",
	//}
	//
	return fmt.Sprintf(baseURL, args...)
}
