package routes

import (
	"nik-api/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRoute(server *gin.Engine, controller controller.IController) {
	api := server.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/submit-ids", controller.SubmitIds)
	v1.POST("/extract/:id", controller.ExtractData)
	v1.POST("/validate", controller.ValidateIds)

	users := v1.Group("/users")
	{
		users.GET("/", controller.GetUsers)
		users.GET("/:id", controller.GetUserById)
		users.GET("/province/:id", controller.GetUsersByProvinceId)
		users.GET("/city/:id", controller.GetUsersByCityId)
		users.GET("/district/:id", controller.GetUsersByDistrictId)
		users.GET("/year/:year", controller.GetUsersByYearOfBirth)
		users.GET("/gender/:gender", controller.GetUsersByGender)
	}

}
