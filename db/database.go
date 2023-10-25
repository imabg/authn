package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var Bun *bun.DB

func CreateDB() *sql.DB {
	var (
		dbname = os.Getenv("DB_NAME")
		dbuser = os.Getenv("DB_USER")
		dbpwd  = os.Getenv("DB_PWD")
		dbhost = os.Getenv("DB_HOST")
		uri    = fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", dbuser, dbpwd, dbhost, dbname)
	)
	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(uri)))
	return db
}

func Init() error {
	sqlDB := CreateDB()
	err := sqlDB.Ping()
	if err != nil {
		return err
	}
	Bun = bun.NewDB(sqlDB, pgdialect.New())
	Bun.AddQueryHook(bundebug.NewQueryHook())
	return nil
}
