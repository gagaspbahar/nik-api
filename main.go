package main

import (
	"nik-api/internal/controller"
	"nik-api/internal/repository"
	"nik-api/internal/routes"
	"nik-api/internal/service"
	db "nik-api/pkg/database"
	"nik-api/pkg/server"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"
)

func main() {
	db := db.NewDb()

	defer db.Close()

	server := server.NewServer()

	repository := repository.NewRepository()
	service := service.NewService(repository, db)
	controller := controller.NewController(service)

	routes.NewRoute(server.App(), controller)

	server.Start(":8080")
}
