package main

import (
	"log"

	"github.com/osami/goomu/internal/data"
)

func main() {
	// Initialize database connection
	data.InitDB()
	log.Println("Database connection established")

	// Example of creating a new user
	user := &data.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "securepassword", // In production, make sure to hash passwords
	}

	result := data.DB.Create(user)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
	}

	// Keep the application running
	select {}
}
