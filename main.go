package main

import (
	"time"
	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000"},
    AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
    AllowHeaders:     []string{"Origin", "Content-Type"},
    ExposeHeaders:    []string{"Content-Length", "Content-Type"},
    AllowCredentials: true,

    MaxAge: 12 * time.Hour,
  }))

	router.POST("/employee", controllers.CreateEmployee)
	router.GET("/employees", controllers.GetEmployees)
	router.GET("/employee/:id", controllers.GetEmployeeByID)
	router.DELETE("/employee/:id", controllers.DeleteEmployee)
	router.PUT("/employee/:id", controllers.UpdateEmployee)
	
	router.Run() // listen and serve on 0.0.0.0:5000
}