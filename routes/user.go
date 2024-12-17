package routes

import (
	"wasatext/handlers"

	"github.com/gorilla/mux"
)

// RegisterUserRoutes registers user-related routes to the router.
func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/user", handlers.UserHandler).Methods("GET")
}
