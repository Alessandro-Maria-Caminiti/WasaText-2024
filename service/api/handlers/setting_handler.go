package handlers

import (
	"encoding/json"
	"net/http"
	"wasatext/service/api/models" // Update this import path
)

var currentSettings = models.Settings{
	Theme:         "dark", // Default theme
	Notifications: true,   // Default notification setting
}

// GetUserSettings handles the GET /settings route
func GetUserSettings(w http.ResponseWriter, r *http.Request) {
	// Return the current settings as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"settings": currentSettings,
	})
}

// UpdateUserSettings handles the POST /settings route
func UpdateUserSettings(w http.ResponseWriter, r *http.Request) {
	var newSettings models.Settings

	// Decode the request body into the newSettings struct
	err := json.NewDecoder(r.Body).Decode(&newSettings)
	if err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Here we can validate the theme (e.g., only "light" or "dark" are allowed)
	if newSettings.Theme != "light" && newSettings.Theme != "dark" {
		http.Error(w, "Invalid theme value provided.", http.StatusBadRequest)
		return
	}

	// Update the global settings with the new settings
	currentSettings = newSettings

	// Respond with the success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "Settings updated successfully",
	})
}
