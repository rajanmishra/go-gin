package main

import (
	"github.com/gin-gonic/gin"

	"rajan/employee/service"
)

func main() {
	router := gin.Default()
	router.GET("/employee", service.GetEmployees)
	router.POST("/employee", service.CreateEmployee)
	router.GET("/employee/:id", service.GetEmployeeByID)
	router.PUT("/employee/:id", service.UpdateEmployee)
	router.DELETE("/employee/:id", service.DeleteEmployee)

	router.Run("localhost:8080")
}
