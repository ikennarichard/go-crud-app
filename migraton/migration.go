package main

import (
	"github.com/ikennarichard/go-crud-app/initializers"
	"github.com/ikennarichard/go-crud-app/models"
)

func main() {
	
	// migrate the schema
	initializers.ConnectDB()
	initializers.DB.AutoMigrate(&models.Employee{})
}