package handlers

import (
	"encoding/json"
	"net/http"
	"wasatext/service/database"
	 
	"github.com/gorilla/mux"
)

// DeleteMessageHandler gestisce la richiesta di eliminazione di un messaggio
func DeleteMessageHandler(database database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	// Estrarre i parametri dall'URL
	vars := mux.Vars(r)
	messageId:= vars["messageId"]
	if messageId == "" {
		respondWithErrordeletemsg(w, http.StatusBadRequest, "ID messaggio non valido")
		return
	}

	convId := vars["sourceChatId"]
	if convId == "" {
		respondWithErrordeletemsg(w, http.StatusBadRequest, "ID conversazione non valido")
		return
	}

	// Tentativo di eliminare il messaggio dal database
	err := database.DeleteMessage(messageId, convId)
	if err != nil {
		respondWithErrordeletemsg(w, http.StatusNotFound, "Messaggio non trovato o errore nell'eliminazione")
		return
	}

	// Risposta con successo
	w.WriteHeader(http.StatusNoContent) // 204 No Content
}}


// utils.respondWithError invia una risposta JSON di errore
func respondWithErrordeletemsg(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
