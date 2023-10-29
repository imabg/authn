package store

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresStore() (*sqlx.DB, error) {
	conn, err := sqlx.Connect("postgres", os.Getenv("DB_URI"))
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}
	return conn, nil
}
