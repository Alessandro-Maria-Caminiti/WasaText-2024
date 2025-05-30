package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"WasaText/service/database"
	"github.com/julienschmidt/httprouter"
)

// Request structure for adding users to a group
type AddToGroupRequest struct {
	GroupName string   `json:"groupName"`
	Names     []string `json:"names"`
}

// Response structure for adding users to a group
type AddToGroupResponse struct {
	Message string `json:"message"`
}

// addToGroup handles adding users to a group (creates group if not existing)
func (rt *_router) addToGroup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set response content type
	w.Header().Set("Content-Type", "application/json")

	// Extract the username from the request context.
	username, ok := r.Context().Value(usernameKey).(string)
	if !ok || username == "" {
		// If the username is missing or invalid, respond with 401 Unauthorized.
		http.Error(w, `{"error": "unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// Parse request body
	var req AddToGroupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `{"error": "invalid request body"}`, http.StatusBadRequest)
		return
	}

	// Validate input
	if len(req.GroupName) < 3 || len(req.GroupName) > 16 {
		http.Error(w, `{"error": "group name must be between 3 and 16 characters"}`, http.StatusBadRequest)
		return
	}

	if len(req.Names) == 0 {
		http.Error(w, `{"error": "at least one other user must be added"}`, http.StatusBadRequest)
		return
	}

	for _, newMember := range req.Names {
		_, err = rt.db.GetUser(newMember)
		if err != nil {
			// Check if error is "user not found"
			if errors.Is(err, database.ErrUserNotFound) {
				http.Error(w, `{"error": "User '`+newMember+`' does not exist"}`, http.StatusBadRequest)
				return
			}
			log.Print("other database error")
			// if different error
			http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
			return
		}
	}

	// Add users to group in database
	err = rt.db.AddToGroup(req.GroupName, req.Names, username)
	log.Print(err)
	if err != nil {
		http.Error(w, `{"error": "failed to add users to group: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	// Send success response
	response := AddToGroupResponse{Message: "Users successfully added to the group. The group was created if it did not exist."}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Handle any potential error during JSON encoding.
		http.Error(w, `{"error": "failed to encode response"}`, http.StatusInternalServerError)
		return
	}
}
