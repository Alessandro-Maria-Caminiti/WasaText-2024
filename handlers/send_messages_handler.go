package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// MessageRequest represents the expected request body for sending a message
type MessageRequest struct {
	SenderID  string    `json:"senderId"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// MessageResponse represents the response sent back to the client
type MessageResponse struct {
	MessageID string `json:"messageId"`
}

// SendMessageHandler handles sending a message to a specific chat
func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Extract sourceChatId from path parameters
	vars := mux.Vars(r)
	sourceChatId := vars["sourceChatId"]

	// Parse and validate request body
	var messageReq MessageRequest
	if err := json.NewDecoder(r.Body).Decode(&messageReq); err != nil {
		log.Printf("Error decoding request body: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid input data"})
		return
	}

	// Validate the senderId
	if messageReq.SenderID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Sender ID is required"})
		return
	}

	// Validate content (emoji, text, or URL)
	if !isValidContent(messageReq.Content) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid content format"})
		return
	}

	// Generate a unique message ID
	messageID := uuid.New().String()

	// Mock saving the message to a database (not implemented)
	log.Printf("Message sent in chat %s by sender %s: %s", sourceChatId, messageReq.SenderID, messageReq.Content)

	// Respond with the generated message ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(MessageResponse{MessageID: messageID})
}

// isValidContent validates message content
func isValidContent(content string) bool {
	// Emoji pattern
	emojiPattern := regexp.MustCompile(`(?:[\x{1F600}-\x{1F64F}]|[\x{1F300}-\x{1F5FF}])+`)

	// URL pattern (simplified)
	urlPattern := regexp.MustCompile(`^https?://\S+$`)

	// Text (any non-empty string is valid)
	if content == "" {
		return false
	}

	return emojiPattern.MatchString(content) || urlPattern.MatchString(content) || len(content) > 0
}
