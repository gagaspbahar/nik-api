package service

import (
	"database/sql"
	"nik-api/internal/repository"
	"nik-api/internal/schema"
	"nik-api/internal/util"

	"github.com/gin-gonic/gin"
)

type IService interface {
	InsertUsers(ctx *gin.Context, users []schema.User) error
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

func (service *service) InsertUsers(ctx *gin.Context, users []schema.User) error {
	tx, err := service.db.Begin()

	defer util.CommitOrRollback(tx)

	if err != nil {
		return err
	}

	err = service.repository.InsertUsers(ctx, tx, users)

	if err != nil {
		return err
	}

	return nil
}
