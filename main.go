package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/go-crud-app/controllers"
	"github.com/ikennarichard/go-crud-app/initializers"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	// Gin
	router := gin.Default()

	router.POST("/employee", controllers.CreateEmployee)
	router.GET("/employees", controllers.GetEmployees)
	router.GET("/employee/:id", controllers.GetEmployeeByID)
	router.DELETE("/employee/:id", controllers.DeleteEmployee)
	router.PUT("/employee/:id", controllers.UpdateEmployee)
	
	router.Run() // listen and serve on 0.0.0.0:5000
}