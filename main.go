package main

import (
	"fmt"
	"log"
	"net/http"
	"orders/config"
	handlers "orders/handler"
	orders "orders/models"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	ordersRepo := orders.NewOrdersRepository(db)
	ordersHandler := &handlers.OrdersHandler{Repo: ordersRepo}

	// routes
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "https://editor.swagger.io")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")
			w.WriteHeader(http.StatusNoContent) // No content in preflight response
			return
		}

		// Handle actual requests
		ordersHandler.ServeHTTP(w, r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://editor.swagger.io"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Accept"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		Debug:            true,
	})

	corsHandler := corsOptions.Handler(http.DefaultServeMux)

	// Start the server
	fmt.Printf("Server started at localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, corsHandler); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
