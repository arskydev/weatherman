package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/arskydev/weatherman/pkg/users"
)

const (
	usersTable = "users"
)

type PostgresAuth struct {
	conn *sql.DB
}

func NewPostgresAuth(conn *sql.DB) *PostgresAuth {
	return &PostgresAuth{
		conn: conn,
	}
}

func (r *PostgresAuth) CreateUser(u users.User) (id int, err error) {
	query := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) values ($1, $2, $3) RETURNING ID", usersTable)

	args := []interface{}{u.Username, u.Email, u.Password}
	row, err := r.queryRow(query, args)

	if err != nil {
		return 0, err
	}

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostgresAuth) GetUser(username, password string) (*users.User, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM  %s WHERE username=$1 AND password_hash=$2", usersTable)

	args := []interface{}{username, password}
	row, err := r.queryRow(query, args)

	if err != nil {
		return nil, err
	}

	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return &users.User{
		ID: id,
	}, nil
}

func (r *PostgresAuth) queryRow(query string, args []interface{}) (*sql.Row, error) {
	c := make(chan *sql.Row)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func(query string, args []interface{}) {
		c <- r.conn.QueryRow(query, args...)
	}(query, args)

	select {
	case <-ctx.Done():
		return nil, errors.New("database query row timeout error")
	case row := <-c:
		return row, nil
	}
}
