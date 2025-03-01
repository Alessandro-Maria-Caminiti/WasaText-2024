package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AddToGroupRequest represents the expected request body
type AddToGroupRequest struct {
	UserID string `json:"userId"`
	Role   string `json:"role"`
}

// AddToGroupResponse represents a successful response
type AddToGroupResponse struct {
	Status string `json:"status"`
}

// AddToGroupHandler handles adding a user to a group chat
func AddToGroupHandler(w http.ResponseWriter, r *http.Request) {
	// Extract sourceChatId from path parameters
	vars := mux.Vars(r)
	sourceChatId := vars["sourceChatId"]

	// Parse and validate request body
	var addReq AddToGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&addReq); err != nil {
		log.Printf("Error decoding request body: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid input data"})
		return
	}

	// Validate userId
	if addReq.UserID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "UserID is required"})
		return
	}

	// Validate role
	if addReq.Role != "member" && addReq.Role != "admin" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid role specified"})
		return
	}

	// Mock adding user to the group (database logic would go here)
	log.Printf("User %s added to group %s as %s", addReq.UserID, sourceChatId, addReq.Role)

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(AddToGroupResponse{
		Status: "User added to group successfully",
	})
}
