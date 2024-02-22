package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/go-crud-app/initializers"
	"github.com/ikennarichard/go-crud-app/models"
)

func CreateEmployee(c *gin.Context) {
	var employee models.Employee

	var body struct {
		// this is the request body
		Name  string `json:"name"`
		Email string `json:"email"`
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

	c.JSON(http.StatusCreated, employee)

}

func GetEmployees(c *gin.Context) {
	var employees []models.Employee
	
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
