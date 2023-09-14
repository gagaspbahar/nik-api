package database

import (
	"database/sql"
	"log"
)

func NewDb() *sql.DB {
	log.Printf("Connecting to database...")
	db, err := sql.Open("sqlite3", "./sql/data-indonesia.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Printf("Connected to database")
	return db
}
