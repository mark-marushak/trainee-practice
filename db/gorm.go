package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() (*gorm.DB, error) {

	connection, err := connect()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: connection,
	}))

	if err != nil {
		return nil, err
	}

	return db, nil
}
