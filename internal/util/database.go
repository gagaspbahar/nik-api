package util

import (
	"database/sql"
	"log"
)

func CommitOrRollback(tx *sql.Tx) {
	if err := recover(); err != nil {
		errorRollback := tx.Rollback()
		log.Fatal(errorRollback)
		panic(err)
	} else if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}
