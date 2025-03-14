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

func enableCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		handler(w, r)
	}
}

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

	// routes with CORS enabled
	http.HandleFunc("/api/ky026", enableCORS(controller.HandleKY026))
	http.HandleFunc("/api/mq2", enableCORS(controller.HandleMQ2))

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
