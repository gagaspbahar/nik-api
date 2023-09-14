package repository

import (
	"database/sql"
	"nik-api/internal/schema"

	"github.com/gin-gonic/gin"
)

type IRepository interface {
	InsertUsers(ctx *gin.Context, tx *sql.Tx, users []schema.UserSchema) error
	GetProvinceById(ctx *gin.Context, tx *sql.Tx, id string) (schema.Province, error)
	GetCityById(ctx *gin.Context, tx *sql.Tx, id string) (schema.City, error)
	GetDistrictById(ctx *gin.Context, tx *sql.Tx, id string) (schema.District, error)
	GetUsersByProvinceId(ctx *gin.Context, tx *sql.Tx, provinceId int) ([]schema.User, error)
	GetUsersByCityId(ctx *gin.Context, tx *sql.Tx, cityId int) ([]schema.User, error)
	GetUsersByDistrictId(ctx *gin.Context, tx *sql.Tx, districtId int) ([]schema.User, error)
	GetUsersByYearOfBirth(ctx *gin.Context, tx *sql.Tx, yearOfBirth string) ([]schema.User, error)
	GetUsersByGender(ctx *gin.Context, tx *sql.Tx, gender string) ([]schema.User, error)
}

type repository struct{}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) GetProvinceById(ctx *gin.Context, tx *sql.Tx, id string) (schema.Province, error) {
	query := `SELECT id, nama FROM t_provinsi WHERE id = ?`
	row := tx.QueryRowContext(ctx, query, id)

	var province schema.Province
	err := row.Scan(&province.Id, &province.Name)

	if err != nil {
		return schema.Province{}, err
	}

	return province, nil
}

func (r *repository) GetCityById(ctx *gin.Context, tx *sql.Tx, id string) (schema.City, error) {
	query := `SELECT id, nama FROM t_kota WHERE id = ?`
	row := tx.QueryRowContext(ctx, query, id)

	var city schema.City
	err := row.Scan(&city.Id, &city.Name)

	if err != nil {
		return schema.City{}, err
	}

	return city, nil
}

func (r *repository) GetDistrictById(ctx *gin.Context, tx *sql.Tx, id string) (schema.District, error) {
	query := `SELECT id, nama FROM t_kecamatan WHERE id = ?`
	row := tx.QueryRowContext(ctx, query, id)

	var district schema.District
	err := row.Scan(&district.Id, &district.Name)

	if err != nil {
		return schema.District{}, err
	}

	return district, nil
}
