package main

import (
	db "nik-api/pkg/database"
	"nik-api/pkg/server"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"
)

func main() {
	db := db.NewDb()

	defer db.Close()

	server := server.NewServer()

	server.Start(":8080")
}
