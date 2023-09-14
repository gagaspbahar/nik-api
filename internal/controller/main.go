package controller

import (
	"database/sql"
	"errors"
	"nik-api/internal/schema"
	"nik-api/internal/service"
	"nik-api/internal/util"
	"strconv"

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
	if err != nil {
		ctx.JSON(500, util.MakeMultipleErrorResponse(500, "Internal Server Error", err))
		return
	}
	ctx.JSON(200, util.MakeResponse(200, "Success", nil, nil))
}

func (c *controller) ExtractData(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := c.service.ExtractData(ctx, id)

	if err != nil {
		ctx.JSON(500, util.MakeMultipleErrorResponse(500, "Internal Server Error", err))
		return
	}

	ctx.JSON(200, util.MakeResponse(200, "success", nil, data))

}

func (c *controller) ValidateIds(ctx *gin.Context) {}

func (c *controller) GetUsers(ctx *gin.Context) {}

func (c *controller) GetUserById(ctx *gin.Context) {}

func (c *controller) GetUsersByProvinceId(ctx *gin.Context) {
	id := ctx.Param("id")
	provinceId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, util.MakeResponse(400, "Bad Request", err, nil))
		return
	}

	users, err := c.service.GetUsersByProvinceId(ctx, provinceId)

	if err == sql.ErrNoRows {
		ctx.JSON(404, util.MakeResponse(404, "Province id not found", nil, nil))
		return
	}

	if err != nil {
		ctx.JSON(500, util.MakeResponse(500, "Internal Server Error", err, nil))
		return
	}

	ctx.JSON(200, util.MakeResponse(200, "success", nil, users))
}

func (c *controller) GetUsersByCityId(ctx *gin.Context) {
	id := ctx.Param("id")
	cityId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, util.MakeResponse(400, "Bad Request", err, nil))
		return
	}

	users, err := c.service.GetUsersByCityId(ctx, cityId)

	if err == sql.ErrNoRows {
		ctx.JSON(404, util.MakeResponse(404, "City id not found", nil, nil))
		return
	}

	if err != nil {
		ctx.JSON(500, util.MakeResponse(500, "Internal Server Error", err, nil))
		return
	}

	ctx.JSON(200, util.MakeResponse(200, "success", nil, users))
}

func (c *controller) GetUsersByDistrictId(ctx *gin.Context) {
	id := ctx.Param("id")
	districtId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, util.MakeResponse(400, "Bad Request", err, nil))
		return
	}

	users, err := c.service.GetUsersByDistrictId(ctx, districtId)

	if err == sql.ErrNoRows {
		ctx.JSON(404, util.MakeResponse(404, "District id not found", nil, nil))
		return
	}

	if err != nil {
		ctx.JSON(500, util.MakeResponse(500, "Internal Server Error", err, nil))
		return
	}

	ctx.JSON(200, util.MakeResponse(200, "success", nil, users))
}

func (c *controller) GetUsersByYearOfBirth(ctx *gin.Context) {
	year := ctx.Param("year")

	users, err := c.service.GetUsersByYearOfBirth(ctx, year)

	if err == sql.ErrNoRows {
		ctx.JSON(404, util.MakeResponse(404, "Year of birth not found", nil, nil))
		return
	}

	if err != nil {
		ctx.JSON(500, util.MakeResponse(500, "Internal Server Error", err, nil))
		return
	}

	ctx.JSON(200, util.MakeResponse(200, "success", nil, users))
}

func (c *controller) GetUsersByGender(ctx *gin.Context) {
	gender := ctx.Param("gender")

	if !(gender == "m" || gender == "f") {
		ctx.JSON(400, util.MakeResponse(400, "Bad Request", errors.New("invalid gender input"), nil))
		return
	}

	users, err := c.service.GetUsersByGender(ctx, gender)

	if err != nil {
		ctx.JSON(500, util.MakeResponse(500, "Internal Server Error", err, nil))
		return
	}

	ctx.JSON(200, util.MakeResponse(200, "success", nil, users))
}
