package main

import (
	"ady-trans-jaya/config"
	"ady-trans-jaya/models"
	"ady-trans-jaya/routes"
	"log"
)

func main() {

	config.ConnectDB()
	r := routes.SetupRouter()
	err := config.DB.AutoMigrate(&models.User{}, &models.Cars{}, &models.Rentals{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migrated successfully.")
	r.Run(":8080")
}
