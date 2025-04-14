package handlers

import (
	"encoding/json"
	"net/http"
	"wasatext/service/database"
	"log"
	"github.com/gorilla/mux"
)

// ForwardMessageRequest represents the request body for forwarding a message
type ForwardMessageRequest struct {
	MessageID         string    	`json:"messageId" validate:"required"`
}

// ForwardMessageResponse represents the response body for a forwarded message
type ForwardMessageResponse struct {
	ForwardedMessageID string 	`json:"forwardedMessageId"`
	Status             string 	`json:"status"`
}

// ForwardMessageHandler gestisce l'inoltro di un messaggio
func ForwardMessageHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { 
	// Estrai i parametri dalla richiesta
	vars := mux.Vars(r)
	sourceChatID := vars["sourceChatId"]
	targetChatID := vars["targetChatId"]
	SenderID := vars["id"]
	log.Printf("Source:%s, Target:%s", sourceChatID, targetChatID)
	if sourceChatID == "" || targetChatID == "" {
		respondWithErrorfrwmsg(w, http.StatusBadRequest, "ID conversazione non valido")
		return
	}

	// Decodifica il corpo della richiesta
	var req ForwardMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		
		respondWithErrorfrwmsg(w, http.StatusBadRequest, "Formato richiesta non valido")
		return
	}

	// Inoltra il messaggio usando il database
	forwardedMessage, err := db.ForwardMessage(req.MessageID, sourceChatID, targetChatID, SenderID)
	if err != nil {
		
		respondWithErrorfrwmsg(w, http.StatusInternalServerError, "Errore nell'inoltro del messaggio")
		return
	}

	// Rispondi con il nuovo messaggio inoltrato
	response := ForwardMessageResponse{
		ForwardedMessageID: forwardedMessage.MessageId, // MessageId è già stringa
		Status:             "Forwarded",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}}
// utils.respondWithError invia una risposta JSON di errore
func respondWithErrorfrwmsg(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
