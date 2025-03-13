package main

import (
	"ApiMulti/src/application"
	"ApiMulti/src/core"
	"ApiMulti/src/infrastructure/controllers"
	"ApiMulti/src/infrastructure/repositories"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	core.InitDB()
	defer core.DB.Close()

	// Initialize RabbitMQ connection
	core.InitRabbitMQ()
	defer core.RabbitMQConn.Close()
	defer core.RabbitMQChannel.Close()

	// Initialize repository
	repository := repositories.NewMySQLRepository(core.DB)

	// Initialize service
	service := application.NewSensorService(repository)

	// Initialize controller
	controller := controllers.NewSensorController(service)

	// Define routes
	http.HandleFunc("/api/ky026", controller.HandleKY026)
	http.HandleFunc("/api/mq2", controller.HandleMQ2)

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
