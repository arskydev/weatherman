package middleware

import "github.com/arskydev/weatherman/pkg/service"

type Middleware struct {
	auth service.Authorization
}

func New(auth service.Authorization) *Middleware {
	return &Middleware{
		auth: auth,
	}
}
