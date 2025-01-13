package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)



// ChatDetails represents the details of a specific chat
type ChatDetails struct {
	Identifier           string    `json:"identifier"`
	Name                 string    `json:"name"`
	Participants         []User    `json:"participants"`
	LastMessage          string    `json:"lastMessage"`
	LastMessageTimestamp time.Time `json:"lastMessageTimestamp"`
	IsGroupChat          bool      `json:"isGroupChat"`
}

// RetrieveChatDetailsHandler retrieves details for a specific chat
func RetrieveChatDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract sourceChatId from path parameters
	vars := mux.Vars(r)
	sourceChatId := vars["sourceChatId"]

	// Mock data: Pretend we have these chat details in a database
	mockChats := map[string]ChatDetails{
		"chat123": {
			Identifier: "chat123",
			Name:       "General Chat",
			Participants: []User{
				{UserID: "user1", Username: "Alice"},
				{UserID: "user2", Username: "Bob"},
			},
			LastMessage:          "Hello everyone!",
			LastMessageTimestamp: time.Now().Add(-1 * time.Hour),
			IsGroupChat:          true,
		},
		"chat456": {
			Identifier: "chat456",
			Name:       "Private Chat with Bob",
			Participants: []User{
				{UserID: "user1", Username: "Alice"},
				{UserID: "user2", Username: "Bob"},
			},
			LastMessage:          "See you tomorrow!",
			LastMessageTimestamp: time.Now().Add(-2 * time.Hour),
			IsGroupChat:          false,
		},
	}

	// Check if the chat exists
	chatDetails, exists := mockChats[sourceChatId]
	if !exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Chat not found",
		})
		return
	}

	// Return the chat details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(chatDetails); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
	}
}
