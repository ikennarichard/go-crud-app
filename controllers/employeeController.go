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
		c.JSON(http.StatusConflict, gin.H{"error": "Error creating employee"})
	}

	c.JSON(http.StatusCreated, employee)

}

func GetEmployees(c *gin.Context) {
	var employees []models.Employee
	
	initializers.DB.Find(&employees)
	c.JSON(http.StatusOK, employees)
}

func GetEmployeeByID(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")

	// loop over the employees and find the
	// one that the param
	if err := initializers.DB.First(&employee, id).Error; err != nil {
    // Handle the case where the record was not found
    c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
    return
	}
	c.JSON(http.StatusOK, employee)
}


func DeleteEmployee(c *gin.Context) {
	var employee models.Employee
	id := c.Param("id")

	// Delete the employee record
	result := initializers.DB.Delete(&employee, id)
	if result.RowsAffected == 0 {
			// Handle the case where the record was not found
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
} 


func UpdateEmployee(c *gin.Context) {
	var employee models.Employee

	var body struct {
		Name string
	}

	c.Bind(&body)
	id := c.Param("id")
	initializers.DB.First(&employee, id)

	initializers.DB.Model(&employee).Update("name", body.Name)
	c.JSON(http.StatusOK, employee)
}