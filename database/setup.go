package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DBHandler struct {
	DB *sql.DB
}

func ConnectToDataBase() (DB *DBHandler, err error) {
	db, err := sql.Open("mysql", "root:12654778@/proto_practice")
	if err != nil {
		return nil, err
	}
	database := &DBHandler{
		DB: db,
	}
	return database, nil
}

