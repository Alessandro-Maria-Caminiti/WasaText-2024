package handlers

import (
	"net/http"
)

// DeleteMessageHandler handles the deletion of a message
func DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Extract parameters from the URL
	sourceChatID := r.URL.Query().Get("sourceChatId")
	messageID := r.URL.Query().Get("messageId")

	if sourceChatID == "" || messageID == "" {
		http.Error(w, "sourceChatId and messageId are required", http.StatusBadRequest)
		return
	}

	// Mock logic to delete a message
	// In a real implementation, you would interact with a database or another service here
	// For now, we'll assume the deletion is successful

	w.WriteHeader(http.StatusNoContent) // 204 No Content indicates successful deletion
}
