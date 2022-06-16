package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/nix-practice")
	if err != nil {
		return nil, err
	}
	return db, nil
}
