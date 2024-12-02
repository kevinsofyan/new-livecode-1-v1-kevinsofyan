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
	http.Handle("/orders/", ordersHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Start the server
	fmt.Printf("Server started at localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

}
