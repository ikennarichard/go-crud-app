package main

import (
	"log"
	"os"

	"github.com/ikennarichard/go-crud-app/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main () {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// migrate the schema
	db.AutoMigrate(&models.Employee{})
}