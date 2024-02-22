package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/go-crud-app/controllers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Gin
	router := gin.Default()

	router.POST("/employee", controllers.CreateEmployee)
	router.GET("/employees", controllers.GetEmployees)
	// router.GET("/employee/:id", controllers.GetEmployeeByID)
	
	router.Run() // listen and serve on 0.0.0.0:5000
}