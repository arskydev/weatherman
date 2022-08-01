package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/arskydev/weatherman/pkg/repository"
	"github.com/arskydev/weatherman/pkg/users"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "fpHJdgsKfmamkfa9aHUAF98safia9r0er"
	tokenTTL   = 72 * time.Hour
	signingKey = "hfuaihAKF21wd31r9wad8hAID9hhd"
)

type tokenClaims struct {
	jwt.StandardClaims
	Id int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(r repository.Authorization) *AuthService {
	return &AuthService{
		repo: r,
	}
}

func (s *AuthService) CreateUser(u users.User) (id int, err error) {
	if u.Username == "" || u.Email == "" || u.Password == "" {
		return 0, errors.New("invalid User. Some fields are missing")
	}

	u.Password = genratePasswordHash(u.Password)

	return s.repo.CreateUser(u)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, genratePasswordHash(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ValidateToken(header string) (*jwt.Token, error) {
	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("invalid token type")
	}

	if len(headerParts[1]) == 0 {
		return nil, errors.New("token is empty")
	}

	token, err := jwt.ParseWithClaims(headerParts[1], &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing token method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		fmt.Println(headerParts[1])
		return nil, err
	}

	_, ok := token.Claims.(*tokenClaims)

	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return token, nil
}

func genratePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
