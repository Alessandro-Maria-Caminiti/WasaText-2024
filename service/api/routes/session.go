package routes

import (
	"wasatext/service/api/handlers" // Import handlers

	"github.com/gorilla/mux"
)

// RegisterSessionRoutes registers session-related routes to the router.
func RegisterSessionRoutes(router *mux.Router) {
	router.HandleFunc("/session", handlers.LoginHandler).Methods("POST")
}

