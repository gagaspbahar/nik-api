package repository

import (
	"database/sql"
	"nik-api/internal/schema"

	"github.com/gin-gonic/gin"
)

func (r *repository) GetUserById(ctx *gin.Context, tx *sql.Tx, id string) (schema.User, error) {
	query := `SELECT id, name FROM users WHERE id = ?`

	row := tx.QueryRowContext(ctx, query, id)

	var user schema.User
	err := row.Scan(&user.Id, &user.Name)

	if err != nil {
		return schema.User{}, err
	}

	return user, nil
}

func (r *repository) InsertUsers(ctx *gin.Context, tx *sql.Tx, users []schema.User) error {
	query := `INSERT INTO users (id, name) VALUES (?, ?)`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	for _, user := range users {
		_, err := stmt.ExecContext(ctx, user.Id, user.Name)
		if err != nil {
			return err
		}
	}

	return nil
}
