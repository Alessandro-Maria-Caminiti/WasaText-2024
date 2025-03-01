package handlers

import (
	"encoding/json"
	"net/http"
)

// UserHandler handles user-specific GET requests.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "User endpoint"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
