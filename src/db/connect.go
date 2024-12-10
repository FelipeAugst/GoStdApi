package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(url string) (*sql.DB, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
