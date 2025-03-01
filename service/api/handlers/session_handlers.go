package handlers

import (
	"encoding/json"
	"net/http"
	"wasatext/service/api/models" // Import models
)

// LoginHandler handles user login and creation.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserRequest

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid input data. Ensure JSON format is correct."})
		return
	}

	// Validate the user input
	if len(user.Name) < 3 || len(user.Name) > 16 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid name format. Name must be between 3 and 16 characters."})
		return
	}

	// Generate a dummy identifier for this example
	response := map[string]string{"identifier": "abcdef012345"}

	// Send the response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
