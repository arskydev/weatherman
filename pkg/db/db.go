package db

import (
	"database/sql"
	"errors"
)

type Config interface {
	CreateConnectString() string
}

type DBConnect interface {
	NewDBConnect(config Config) error
	GetConn() *sql.DB
	GetDbms() string
	Close() error
}

func NewConnect(cfg Config) (DBConnect, error) {
	switch cfg.(type) {

	case *PGConfig:
		dbms := "postgres"
		db := NewPostgresConnect(dbms)
		if err := db.NewDBConnect(cfg); err != nil {
			return nil, err
		}
		return db, nil

	default:
		return nil, errors.New("DBMS not supported")
	}
}
