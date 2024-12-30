package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Init(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	return db, nil
}
