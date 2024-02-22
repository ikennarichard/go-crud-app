package main

import (
	"github.com/ikennarichard/go-crud-app/initializers"
	"github.com/ikennarichard/go-crud-app/models"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	
	// migrate the schema
	initializers.DB.AutoMigrate(&models.Employee{})
}