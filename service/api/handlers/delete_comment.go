package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// DeleteCommentHandler handles the deletion of a comment on a specific message
func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Extract sourceChatId and commentId from the URL path
	vars := mux.Vars(r)
	sourceChatId := vars["sourceChatId"]
	commentId := vars["commentId"]

	// Simulate checking if the comment exists (you would check a database or in-memory store here)
	commentExists := true // Mock value. Replace with actual comment existence check.
	if !commentExists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Comment not found"})
		return
	}

	// Simulate deleting the comment (database logic would go here)
	log.Printf("Comment %s deleted from chat %s", commentId, sourceChatId)

	// Respond with a 204 status indicating successful deletion
	w.WriteHeader(http.StatusNoContent)
}
