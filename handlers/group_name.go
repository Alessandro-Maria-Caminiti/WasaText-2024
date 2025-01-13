package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ChangeGroupNameHandler handles changing the group name
func ChangeGroupNameHandler(w http.ResponseWriter, r *http.Request) {
	// Extract sourceChatId from the URL path
	vars := mux.Vars(r)
	sourceChatId := vars["sourceChatId"]

	// Parse the JSON body
	var requestData struct {
		Content string `json:"content"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Validate the new group name
	if len(requestData.Content) == 0 {
		http.Error(w, "Invalid name format", http.StatusBadRequest)
		return
	}

	// Simulate updating the group name (in a real scenario, update in database)
	log.Printf("Group %s changed name to: %s", sourceChatId, requestData.Content)

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "Group name updated successfully",
	})
}
