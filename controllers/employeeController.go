package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/go-crud-app/initializers"
	"github.com/ikennarichard/go-crud-app/models"
)

var employees []models.Employee

func CreateEmployee(c *gin.Context) {
	var employee models.Employee

	var body struct {
		// this is the request body
		Name  string `json:"title"`
		Email string `json:"artist"`
	}

	// bind data to context to save it
	c.Bind(&body)

	// create employee
	employee = models.Employee{Name: body.Name, Email: body.Email}

	result := initializers.DB.Create(&employee)

	// check for errors
	if result.Error != nil {
		log.Fatal("Error creating new employee")
	}

}

func GetEmployees(c *gin.Context) {
	initializers.DB.Find(&employees)
	c.JSON(http.StatusOK, employees)
}

// func GetEmployeeByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// loop over the albums and find the
// 	// one that the param
// 	for _, a:= range employees {
// 		if a.ID == id {
// 			c.JSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"message": "Album not found"})
// }
