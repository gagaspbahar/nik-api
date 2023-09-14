package migrate

import (
	db "nik-api/pkg/database"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
)

func migrateUp() {
	db := db.NewDb()
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrate/migrations", "sqlite3", driver)

	if err != nil {
		panic(err)
	}

	m.Up()
}

func migrateDown() {
	db := db.NewDb()
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrate/migrations", "sqlite3", driver)

	if err != nil {
		panic(err)
	}

	m.Down()
}

func migrateNew(string name) {
	db := db.NewDb()
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrate/migrations", "sqlite3", driver)

	if err != nil {
		panic(err)
	}
}
