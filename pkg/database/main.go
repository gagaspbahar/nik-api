package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewDb() *sql.DB {
	log.Printf("Connecting to database...")
	db, err := sql.Open("sqlite3", "./data-wilayah.db")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to database")
	return db
}
