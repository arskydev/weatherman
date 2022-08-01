package repository

import (
	"database/sql"
	"errors"

	"github.com/arskydev/weatherman/pkg/db"
	"github.com/arskydev/weatherman/pkg/users"
)

type Authorization interface {
	CreateUser(u users.User) (id int, err error)
	GetUser(username, password string) (u *users.User, err error)
}

func NewAuthorization(dbmsAuth string, connAuth *sql.DB) (Authorization, error) {
	if dbmsAuth == "" || connAuth == nil {
		return nil, errors.New("empty fileds passed in NewAuthorization()")
	}
	switch dbmsAuth {
	case "postgres":
		return NewPostgresAuth(connAuth), nil
	default:
		return nil, errors.New("DBMS not supported")
	}
}

type Repository struct {
	Authorization
}

func NewRepository(dbAuth db.DBConnect) (*Repository, error) {
	dbmsAuth := dbAuth.GetDbms()
	connAuth := dbAuth.GetConn()
	auth, err := NewAuthorization(dbmsAuth, connAuth)

	if err != nil {
		return nil, err
	}

	return &Repository{
		Authorization: auth,
	}, nil
}
