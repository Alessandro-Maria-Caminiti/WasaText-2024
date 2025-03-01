package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// User represents a participant in a chat
type User struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
}

// Conversation represents a chat or group conversation
type Conversation struct {
	Identifier           string    `json:"identifier"`
	Name                 string    `json:"name"`
	Participants         []User    `json:"participants"`
	LastMessage          string    `json:"lastMessage"`
	LastMessageTimestamp time.Time `json:"lastMessageTimestamp"`
	IsGroupChat          bool      `json:"isGroupChat"`
}

// ConversationsResponse represents the response structure for chats
type ConversationsResponse struct {
	Conversations []Conversation `json:"conversations"`
}

// RetrieveChatsHandler handles the retrieval of user conversations
func RetrieveChatsHandler(w http.ResponseWriter, r *http.Request) {
	// Mocking a list of conversations
	mockConversations := []Conversation{
		{
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
		{
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

	// Prepare response
	response := ConversationsResponse{
		Conversations: mockConversations,
	}

	// Set headers and send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
		return
	}
}
