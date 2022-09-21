package db

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type PGConfig struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string
	DBName   string `yaml:"DBName"`
	SSLMode  string `yaml:"SSLMode"`
}

type PostgresConnect struct {
	dbms string
	conn *sql.DB
}

func NewPGConfig(confPath, dbPass string) (*PGConfig, error) {
	yfile, err := ioutil.ReadFile(confPath)

	if err != nil {
		return nil, err
	}

	conf := &PGConfig{}
	err = yaml.Unmarshal(yfile, conf)

	if err != nil {
		return nil, err
	}

	conf.Password = dbPass

	if conf.Password == "" {
		return nil, errors.New("no password for db set")
	}

	return conf, nil
}

func (cfg *PGConfig) CreateConnectString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
}

func NewPostgresConnect(dbms string) *PostgresConnect {
	return &PostgresConnect{
		dbms: dbms,
	}
}

func (db *PostgresConnect) NewDBConnect(config Config) error {
	conn, err := sql.Open(db.dbms, config.CreateConnectString())

	if err != nil {
		return err
	}

	err = conn.Ping()

	if err != nil {
		return err
	}

	db.conn = conn

	return nil
}

func (db *PostgresConnect) GetConn() *sql.DB {
	return db.conn
}

func (db *PostgresConnect) GetDbms() string {
	return db.dbms
}

func (db *PostgresConnect) Close() error {
	return db.conn.Close()
}
