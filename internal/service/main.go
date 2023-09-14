package service

import (
	"database/sql"
	"errors"
	"nik-api/internal/repository"
	"nik-api/internal/schema"
	"nik-api/internal/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IService interface {
	InsertUsers(ctx *gin.Context, users []schema.User) []error
	ValidateId(ctx *gin.Context, id string) bool
	GetUsersByProvinceId(ctx *gin.Context, provinceId int) ([]schema.User, error)
	GetUsersByCityId(ctx *gin.Context, cityId int) ([]schema.User, error)
	GetUsersByDistrictId(ctx *gin.Context, districtId int) ([]schema.User, error)
	GetUsersByYearOfBirth(ctx *gin.Context, yearOfBirth string) ([]schema.User, error)
	GetUsersByGender(ctx *gin.Context, gender string) ([]schema.User, error)
	ExtractData(ctx *gin.Context, id string) (schema.ExtractResponse, []error)
}

type service struct {
	repository repository.IRepository
	db         *sql.DB
}

func NewService(repo repository.IRepository, db *sql.DB) *service {
	return &service{
		repository: repo,
		db:         db,
	}
}

func (service *service) InsertUsers(ctx *gin.Context, users []schema.User) []error {
	tx, err := service.db.Begin()
	var errs []error
	defer util.CommitOrRollback(tx)

	if err != nil {
		return []error{err}
	}

	for i := 0; i < len(users); i++ {
		valid := service.ValidateId(ctx, users[i].Id)
		if !valid {
			errs = append(errs, errors.New("invalid id "+users[i].Id))
			users = append(users[:i], users[i+1:]...)
			continue
		}
	}

	if len(errs) > 0 {
		return errs
	}

	var userSchemas []schema.UserSchema

	for _, user := range users {
		gender, err := util.GetGenderFromId(user.Id)
		if err != nil {
			return []error{err}
		}
		dateOfBirth, err := util.GetDateOfBirthFromId(user.Id)
		if err != nil {
			return []error{err}
		}
		userSchemas = append(userSchemas, schema.UserSchema{
			Id:          user.Id,
			Name:        user.Name,
			ProvinceId:  user.Id[0:2],
			CityId:      user.Id[0:4],
			DistrictId:  user.Id[0:6],
			Gender:      gender,
			DateOfBirth: dateOfBirth,
		})
	}

	err = service.repository.InsertUsers(ctx, tx, userSchemas)

	if err != nil {
		return []error{err}
	}

	return nil
}

func (service *service) ValidateId(ctx *gin.Context, id string) bool {
	tx, err := service.db.Begin()

	if err != nil {
		return false
	}

	defer util.CommitOrRollback(tx)

	// Get 6 digit of id

	districtId, err := strconv.Atoi(id[0:6])
	if err != nil {
		return false
	}

	// Validate District
	_, err = service.repository.GetDistrictById(ctx, tx, strconv.Itoa(districtId))

	if err != nil {
		return false
	}

	genderId, err := strconv.Atoi(id[6:8])
	if err != nil {
		return false
	}

	// Validate Gender + YoB
	if !((genderId > 0 && genderId < 32) || (genderId > 40 && genderId < 72)) {
		return false
	}

	return true
}

func (service *service) GetUsersByProvinceId(ctx *gin.Context, provinceId int) ([]schema.User, error) {
	tx, err := service.db.Begin()

	if err != nil {
		return []schema.User{}, err
	}

	defer util.CommitOrRollback(tx)

	_, err = service.repository.GetProvinceById(ctx, tx, strconv.Itoa(provinceId))

	if err != nil {
		return []schema.User{}, sql.ErrNoRows
	}

	users, err := service.repository.GetUsersByProvinceId(ctx, tx, provinceId)

	if err != nil {
		return []schema.User{}, err
	}

	if len(users) == 0 {
		return []schema.User{}, errors.New("no user found")
	}

	return users, nil
}

func (service *service) GetUsersByCityId(ctx *gin.Context, cityId int) ([]schema.User, error) {
	tx, err := service.db.Begin()

	if err != nil {
		return []schema.User{}, err
	}

	defer util.CommitOrRollback(tx)

	_, err = service.repository.GetCityById(ctx, tx, strconv.Itoa(cityId))

	if err != nil {
		return []schema.User{}, sql.ErrNoRows
	}

	users, err := service.repository.GetUsersByCityId(ctx, tx, cityId)

	if err != nil {
		return []schema.User{}, err
	}

	if len(users) == 0 {
		return []schema.User{}, errors.New("no user found")
	}

	return users, nil
}

func (service *service) GetUsersByDistrictId(ctx *gin.Context, districtId int) ([]schema.User, error) {
	tx, err := service.db.Begin()

	if err != nil {
		return []schema.User{}, err
	}

	defer util.CommitOrRollback(tx)

	_, err = service.repository.GetDistrictById(ctx, tx, strconv.Itoa(districtId))

	if err != nil {
		return []schema.User{}, sql.ErrNoRows
	}

	users, err := service.repository.GetUsersByDistrictId(ctx, tx, districtId)

	if err != nil {
		return []schema.User{}, err
	}

	if len(users) == 0 {
		return []schema.User{}, errors.New("no user found")
	}

	return users, nil
}

func (service *service) GetUsersByYearOfBirth(ctx *gin.Context, yearOfBirth string) ([]schema.User, error) {
	tx, err := service.db.Begin()

	if err != nil {
		return []schema.User{}, err
	}

	defer util.CommitOrRollback(tx)

	users, err := service.repository.GetUsersByYearOfBirth(ctx, tx, yearOfBirth)

	if err != nil {
		return []schema.User{}, err
	}

	return users, nil
}

func (service *service) GetUsersByGender(ctx *gin.Context, gender string) ([]schema.User, error) {
	tx, err := service.db.Begin()

	if err != nil {
		return []schema.User{}, err
	}

	defer util.CommitOrRollback(tx)

	users, err := service.repository.GetUsersByGender(ctx, tx, gender)

	if err != nil {
		return []schema.User{}, err
	}

	return users, nil
}

func (service *service) ExtractData(ctx *gin.Context, id string) (schema.ExtractResponse, []error) {
	var errs []error

	tx, err := service.db.Begin()

	if err != nil {
		return schema.ExtractResponse{}, []error{}
	}

	defer util.CommitOrRollback(tx)

	province, err := service.repository.GetProvinceById(ctx, tx, id[0:2])

	if err != nil {
		errs = append(errs, errors.New("invalid province id"))
	}

	city, err := service.repository.GetCityById(ctx, tx, id[0:4])

	if err != nil {
		errs = append(errs, errors.New("invalid city id"))
	}

	district, err := service.repository.GetDistrictById(ctx, tx, id[0:6])

	if err != nil {
		errs = append(errs, errors.New("invalid district id"))
	}

	dob, err := util.GetDateOfBirthFromId(id)

	if err != nil {
		errs = append(errs, errors.New("invalid date of birth"))
	}

	gender, err := util.GetGenderFromId(id)

	if err != nil {
		errs = append(errs, errors.New("invalid gender"))
	}

	if len(errs) > 0 {
		return schema.ExtractResponse{}, errs
	}

	return schema.ExtractResponse{
		Province: province.Name,
		City:     city.Name,
		District: district.Name,
		Id:       id,
		Dob:      dob,
		Gender:   gender,
	}, nil

}
