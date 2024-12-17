package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"wasatext/routes" // Import routes package
)


func main() {
	// Initialize router
	router := mux.NewRouter()

	// Load routes
	routes.RegisterSessionRoutes(router) // Session routes
	routes.RegisterUserRoutes(router)    // User routes

	// Start the server
	log.Println("Server is listening on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}