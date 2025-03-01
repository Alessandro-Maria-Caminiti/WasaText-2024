package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// LeaveGroupChatHandler handles the logic for a user to leave a group chat
func LeaveGroupChatHandler(w http.ResponseWriter, r *http.Request) {
	// Extract sourceChatId from the URL path
	vars := mux.Vars(r)
	sourceChatId := vars["sourceChatId"]

	// Simulate checking if the group exists (replace with actual group check logic)
	groupExists := true // Mock value, you would use a database or store check here
	if !groupExists {
		// Respond with a 404 if the group doesn't exist
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Group not found"})
		return
	}

	// Simulate the logic of leaving the group (actual removal logic would be here)
	log.Printf("User left group %s", sourceChatId)

	// Respond with a 204 No Content status for successful operation
	w.WriteHeader(http.StatusNoContent)
}
