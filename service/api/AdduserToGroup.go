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
type AddUserToGroupRequest struct {
	Name string `json:"name"`
}

// addUserToGroup handles the HTTP request to add a user to a group
func (rt *_router) addUserToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Set response content type
	w.Header().Set("Content-Type", "application/json")
	// Get old groupname from Path parameter
	GroupName := ps.ByName("groupname")
	if GroupName == "" {
		http.Error(w, `{"error": "missing groupname parameter"}`, http.StatusBadRequest)
		return
	}

	// Parse request body
	var req AddUserToGroupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `{"error": "invalid request body"}`, http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, `{"error": "at least one user must be written"}`, http.StatusBadRequest)
		return
	}
	_, err = rt.db.GetUser(req.Name)
	if err != nil {
		// Check if error is "user not found"
		if errors.Is(err, database.ErrUserNotFound) {
			http.Error(w, `{"error": "User '`+req.Name+`' does not exist"}`, http.StatusBadRequest)
			return
		}
		log.Print("other database error")
		// if different error
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	// Add users to group in database
	err = rt.db.AddUserToGroup(GroupName, req.Name)
	if err != nil {
		log.Print(err)
		http.Error(w, `{"error": "failed to add users to group: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
}
