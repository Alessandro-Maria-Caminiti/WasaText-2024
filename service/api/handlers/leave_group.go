package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"wasatext/service/database"
	"github.com/gorilla/mux"
)


// LeaveGroupChatHandler gestisce la rimozione di un utente da un gruppo
func LeaveGroupChatHandler(database database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	// Estrai gli ID dalla richiesta
	vars := mux.Vars(r)
	sourceChatId := vars["sourceChatId"] // ID del gruppo
	userId := vars["id"] // ID dell'utente che vuole uscire
	log.Printf("UserId = %s, SourceChatId = %s", userId, sourceChatId)
	// Controllo che userId sia stato fornito
	if userId == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "User ID is required"})
		return
	}

	// Controlla se il gruppo esiste
	groupExists, err := database.GroupExists(sourceChatId)
	if err != nil {
		log.Printf("Errore nel controllo del gruppo: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
		return
	}

	if !groupExists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Group not found"})
		return
	}

	// Controlla se l'utente è un membro del gruppo
	isMember, err := database.CheckMember(userId, sourceChatId)
	if err != nil {
		log.Printf("Errore nel controllo dell'utente nel gruppo: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
		return
	}

	if !isMember {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"error": "User is not a member of the group"})
		return
	}

	// Rimuove l'utente dal gruppo
	err = database.DeleteUserFromGroup(userId, sourceChatId)
	if err != nil {
		log.Printf("Errore nella rimozione dell'utente dal gruppo: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to leave group"})
		return
	}

	// Log e risposta di successo
	log.Printf("Utente %s rimosso dal gruppo %s con successo", userId, sourceChatId)
	w.WriteHeader(http.StatusNoContent) // 204 No Content per successo
}}