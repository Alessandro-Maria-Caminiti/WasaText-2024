package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

// ForwardMessageRequest represents the request body for forwarding a message
type ForwardMessageRequest struct {
	MessageID         string    `json:"messageId" validate:"required"`
	SenderID          string    `json:"senderId" validate:"required"`
	OriginalTimestamp time.Time `json:"originalTimestamp" validate:"required"`
}

// ForwardMessageResponse represents the response body for a forwarded message
type ForwardMessageResponse struct {
	ForwardedMessageID string `json:"forwardedMessageId"`
	Status             string `json:"status"`
}

// ForwardMessageHandler handles the forwarding of a message
func ForwardMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Extract parameters from the URL
	sourceChatID := r.URL.Query().Get("sourceChatId")
	targetChatID := r.URL.Query().Get("targetChatId")

	if sourceChatID == "" || targetChatID == "" {
		http.Error(w, "sourceChatId and targetChatId are required", http.StatusBadRequest)
		return
	}

	// Parse request body
	var req ForwardMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate request
	if req.MessageID == "" || req.SenderID == "" {
		http.Error(w, "messageId and senderId are required", http.StatusBadRequest)
		return
	}

	// Mock logic for forwarding a message
	response := ForwardMessageResponse{
		ForwardedMessageID: "forwardedMsg101", // This would be dynamically generated in real logic
		Status:             "Forwarded",
	}

	// Respond with the forwarded message details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
