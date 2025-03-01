package routes

import (
	"wasatext/service/api/handlers"

	"github.com/gorilla/mux"
)

// RegisterUserRoutes registers user-related routes to the router.
func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/user", handlers.UserHandler).Methods("GET")
}
