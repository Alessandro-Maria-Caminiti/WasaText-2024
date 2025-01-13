package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"wasatext/models" // Import models
)

// SetUsernameHandler handles the /session/setusername POST request.
func SetUsernameHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserRequest

	// Decode the request body into UserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid username format. Ensure JSON format is correct."})
		return
	}

	// Validate the username format and length
	if err := validateUsername(user.Name); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Check if the username already exists (this is just a placeholder; use real DB logic in production)
	if usernameExists(user.Name) {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"error": "The username is already taken. Please choose another."})
		return
	}

	// Simulate the process of setting the username (e.g., saving it in a DB)
	// For now, return the username as the identifier.
	response := map[string]string{"identifier": user.Name}

	// Send success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// validateUsername validates the username against the given pattern and length
func validateUsername(username string) error {
	// Check length
	if len(username) < 3 || len(username) > 32 {
		return fmt.Errorf("Username must be between 3 and 32 characters long.")
	}

	// Check for valid characters using regular expression
	// This regex allows only alphanumeric characters, dots, and underscores.
	var validUsernamePattern = `^[a-zA-Z0-9._]+$`
	matched, _ := regexp.MatchString(validUsernamePattern, username)
	if !matched {
		return fmt.Errorf("Username contains invalid characters or does not meet length requirements.")
	}

	return nil
}

// usernameExists is a placeholder function to simulate checking if a username exists in a database.
func usernameExists(username string) bool {
	// This should interact with your database to check if the username already exists.
	// For now, we assume no username exists.
	return false
}
