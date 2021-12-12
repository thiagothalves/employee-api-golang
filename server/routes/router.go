package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagothalves/employee-api-golang/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		employees := main.Group("employees")
		{
			employees.GET("/", controllers.FindAllEmployees)
		}
	}

	return router
}
