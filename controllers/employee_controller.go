package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagothalves/employee-api-golang/core"
)

func FindAllEmployees(c *gin.Context) {

	employees := [2]core.Employee{
		{
			Id:     1,
			Name:   "Thiago",
			Office: "Developer",
		},
		{
			Id:     2,
			Name:   "Marcela",
			Office: "Developer",
		},
	}

	c.JSON(200, employees)
}
