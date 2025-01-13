package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CommentRequest represents the structure of the comment request body
type CommentRequest struct {
	SenderID string `json:"senderId"`
	Content  string `json:"content"`
}

// CommentResponse represents the structure of the response body after commenting
type CommentResponse struct {
	MessageID string `json:"messageId"`
	Status    string `json:"status"`
}

// CommentMessageHandler handles posting a comment on a specific message
func CommentMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Extract sourceChatId from the URL path
	vars := mux.Vars(r)
	sourceChatId := vars["sourceChatId"]

	// Parse the request body
	var commentReq CommentRequest
	if err := json.NewDecoder(r.Body).Decode(&commentReq); err != nil {
		log.Printf("Error decoding request body: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid input data"})
		return
	}

	// Validate senderId and content
	if commentReq.SenderID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Sender ID is required"})
		return
	}

	if len(commentReq.Content) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Content cannot be empty"})
		return
	}

	// Simulate checking if the message exists in the chat (you would check a database or in-memory store here)
	messageExists := true // Mock value. Replace with actual message existence check.
	if !messageExists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Message not found"})
		return
	}

	// Mock adding the comment to the message (database logic would go here)
	log.Printf("Comment by %s on message in chat %s: %s", commentReq.SenderID, sourceChatId, commentReq.Content)

	// Respond with the message ID and status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CommentResponse{
		MessageID: "msg789", // Mocked message ID
		Status:    "Sent",   // Success status
	})
}
