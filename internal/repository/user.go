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

func (r *repository) InsertUsers(ctx *gin.Context, tx *sql.Tx, users []schema.UserSchema) error {
	query := `INSERT INTO users (id, name, district_id, city_id, province_id, gender, date_of_birth) VALUES (?, ?, ?, ?, ?, ?, ?)`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	for _, user := range users {
		_, err := stmt.ExecContext(ctx, user.Id, user.Name, user.DistrictId, user.CityId, user.ProvinceId, user.Gender, user.DateOfBirth)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) GetUsersByProvinceId(ctx *gin.Context, tx *sql.Tx, provinceId int) ([]schema.User, error) {
	query := `SELECT u.id, u.name FROM users u WHERE u.province_id = ?`

	rows, err := tx.QueryContext(ctx, query, provinceId)
	if err != nil {
		return nil, err
	}

	var users []schema.User
	for rows.Next() {
		var user schema.User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *repository) GetUsersByCityId(ctx *gin.Context, tx *sql.Tx, cityId int) ([]schema.User, error) {
	query := `SELECT u.id, u.name FROM users u WHERE u.city_id = ?`

	rows, err := tx.QueryContext(ctx, query, cityId)
	if err != nil {
		return nil, err
	}

	var users []schema.User
	for rows.Next() {
		var user schema.User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *repository) GetUsersByDistrictId(ctx *gin.Context, tx *sql.Tx, districtId int) ([]schema.User, error) {
	query := `SELECT u.id, u.name FROM users u WHERE u.district_id = ?`

	rows, err := tx.QueryContext(ctx, query, districtId)
	if err != nil {
		return nil, err
	}

	var users []schema.User
	for rows.Next() {
		var user schema.User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *repository) GetUsersByYearOfBirth(ctx *gin.Context, tx *sql.Tx, yearOfBirth string) ([]schema.User, error) {
	query := `SELECT u.id, u.name FROM users u WHERE date_of_birth LIKE ? || '%'`

	rows, err := tx.QueryContext(ctx, query, yearOfBirth)
	if err != nil {
		return nil, err
	}

	var users []schema.User
	for rows.Next() {
		var user schema.User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *repository) GetUsersByGender(ctx *gin.Context, tx *sql.Tx, gender string) ([]schema.User, error) {
	query := `SELECT u.id, u.name FROM users u WHERE u.gender = ?`

	rows, err := tx.QueryContext(ctx, query, gender)
	if err != nil {
		return nil, err
	}

	var users []schema.User
	for rows.Next() {
		var user schema.User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
