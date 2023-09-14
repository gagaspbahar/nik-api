package controller

import (
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
	var users []schema.User

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err := c.service.InsertUsers(ctx, users)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Success",
	})
}
