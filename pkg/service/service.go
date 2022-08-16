package service

import (
	"github.com/arskydev/weatherman/pkg/repository"
	"github.com/arskydev/weatherman/pkg/users"
	"github.com/dgrijalva/jwt-go"
)

type Authorization interface {
	CreateUser(u users.User) (id int, err error)
	GenerateToken(username, password string) (token string, err error)
	ValidateToken(string) (*jwt.Token, error)
}

type Service struct {
	Authorization Authorization
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r.Authorization),
	}
}
