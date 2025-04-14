package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"wasatext/service/database" 
	"github.com/gorilla/mux"
)

// ChangeGroupNameHandler handles changing the group name
func ChangeGroupNameHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    // Extract sourceChatId from the URL path
    vars := mux.Vars(r)
    sourceChatId := vars["sourceChatId"]

    // Parse the JSON body
    var requestData struct {
        GroupName string `json:"groupName"`
    }
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, "Invalid input data", http.StatusBadRequest)
        return
    }

    // Validate the new group name
    if len(requestData.GroupName) == 0 {
        http.Error(w, "Invalid name format", http.StatusBadRequest)
        return
    }


    // Update group name in database
    err := db.SetGroupName(sourceChatId, requestData.GroupName) // ✅ CORRETTO: usiamo "=" invece di ":="
    if err != nil {
        log.Printf("Error updating group name: %v", err)
        http.Error(w, "Failed to update group name", http.StatusInternalServerError)
        return
    }

    // Respond with success
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "status": "Group name updated successfully",
    })
}}