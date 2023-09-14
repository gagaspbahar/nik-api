package controller

import (
	"log"
	"nik-api/internal/schema"
	"nik-api/internal/service"

	"github.com/gin-gonic/gin"
)

type IController interface {
	SubmitIds(ctx *gin.Context)
	ExtractData(ctx *gin.Context)
	ValidateIds(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	GetUsersByProvinceId(ctx *gin.Context)
	GetUsersByCityId(ctx *gin.Context)
	GetUsersByDistrictId(ctx *gin.Context)
	GetUsersByYearOfBirth(ctx *gin.Context)
	GetUsersByGender(ctx *gin.Context)
}

type controller struct {
	service service.IService
}

func NewController(service service.IService) *controller {
	return &controller{
		service,
	}
}

func (c *controller) SubmitIds(ctx *gin.Context) {
	var users schema.SubmitIdsRequest

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err := c.service.InsertUsers(ctx, users.Data)
	log.Println(err)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Success",
	})
}

func (c *controller) ExtractData(ctx *gin.Context) {}

func (c *controller) ValidateIds(ctx *gin.Context) {}

func (c *controller) GetUsers(ctx *gin.Context) {}

func (c *controller) GetUserById(ctx *gin.Context) {}

func (c *controller) GetUsersByProvinceId(ctx *gin.Context) {}

func (c *controller) GetUsersByCityId(ctx *gin.Context) {}

func (c *controller) GetUsersByDistrictId(ctx *gin.Context) {}

func (c *controller) GetUsersByYearOfBirth(ctx *gin.Context) {}

func (c *controller) GetUsersByGender(ctx *gin.Context) {}
