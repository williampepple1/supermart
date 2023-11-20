package main

import (
	"supermart/config"
	"supermart/models"
	"supermart/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, err := config.InitDB() // Initialize the database connection
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	// Migrate the schema

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Branch{})
	db.AutoMigrate(&models.Inventory{})
	db.AutoMigrate(&models.Sale{})
	db.AutoMigrate(&models.Product{})

	// Set up routes
	routes.SetupUserRoutes(r, db) // Setup user routes

	r.Run(":8080")
}
