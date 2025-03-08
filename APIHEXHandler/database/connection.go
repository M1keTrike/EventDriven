package database

import (
	"database/sql"
	"log"
)

func Connect() (*sql.DB, error) {
	LoadConfig()
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}
	return db, nil
}
