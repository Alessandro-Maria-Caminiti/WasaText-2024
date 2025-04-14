package handlers

import (
	"encoding/json"
	"net/http"
	"wasatext/service/database"
	"github.com/gorilla/mux"
	"log"
	"wasatext/service/api/structs"
	"time"
	"wasatext/service/api/images"
)
 
type Group struct {
    Name    string `json:"name"`
}

// CreateGroupHandler gestisce la creazione di un nuovo gruppo
func CreateGroupHandler(database database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	var group Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		log.Printf("Dati non validi:", err.Error())
		respondWithErrorgrp(w, http.StatusBadRequest, "Dati non validi")
		return
	}
	vars := mux.Vars(r)
	userId := vars["id"]

	newGroup, err := database.CreateGroup(group.Name, userId)
	if err != nil {
		log.Printf("Crezione del gruppo errore:", err.Error())
		respondWithErrorgrp(w, http.StatusInternalServerError, "Errore nella creazione del gruppo")
		return
	}

	users, err := database.GetUserByName("SERVER")
	if err != nil {
		log.Printf("Error searching users: %v", err)
		http.Error(w, "Error searching users", http.StatusInternalServerError)
		return
	}

	log.Printf(newGroup.ConversationId)

	err = database.AddUserConv(newGroup.ConversationId, users.UserId)
	if err != nil {
		http.Error(w, "Errore nell'aggiunta dell'utente alla conversazione", http.StatusInternalServerError)
		return
	}

	msg := structs.Message{
		SenderUserId:   "00000000-0000-0000-0000-000000000000",
		ConversationId: newGroup.ConversationId,
		Text:           "Conversation Created",
		Photo:          "", // Se il messaggio contiene un file, salviamo il percorso
		SendTime:       time.Now(),
		Status:         "sent", // Imposta lo stato del messaggio
	}
	_, err = database.CreateMessage(msg)
	if err != nil {
		log.Printf("Errore nella creazione del messaggio: %v", err)
		respondWithErrorsndmsg(w, http.StatusInternalServerError, "Errore nell'invio del messaggio")
		return
	}
	
	newGroup.GrouPhoto = images.SetDefaultPhotoGroup(newGroup.GroupId)
	respondWithJSON(w, http.StatusCreated, newGroup)
}}

// GetGroupHandler restituisce un gruppo dato il suo ID
func GetGroupHandler(database database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	GroupId := vars["SourceChatId"]
	group, err := database.GetGroupById(GroupId)
	if err != nil {
		respondWithErrorgrp(w, http.StatusNotFound, "Gruppo non trovato")
		return
	}

	respondWithJSON(w, http.StatusOK, group)
}}

// DeleteGroupHandler gestisce l'eliminazione di un gruppo
func DeleteGroupHandler(database database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId := vars["SourceChatId"]
	convId, err := database.GetConvGroup(groupId)
	if err != nil {
		respondWithErrorgrp(w, http.StatusNotFound, "Conversazione non trovata per questo gruppo")
		return
	}
	if err := database.DeleteGroup(convId); err != nil {
		respondWithErrorgrp(w, http.StatusInternalServerError, "Errore nella cancellazione del gruppo")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Gruppo eliminato con successo"})
}}

// GetConvGroupHandler ottiene l'ID della conversazione associata a un gruppo
func GetConvGroupHandler(database database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId := vars["SourceChatId"]
	convId, err := database.GetConvGroup(groupId)
	if err != nil {
		respondWithErrorgrp(w, http.StatusNotFound, "Conversazione non trovata per questo gruppo")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"conversationId": convId})
}}

// UpdateGroupIdHandler aggiorna l'ID di un gruppo
func UpdateGroupIdHandler(database database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	groupId := r.FormValue("groupId")
	conversationId := r.FormValue("conversationId")


	if err := database.UpdateGroupId(groupId, conversationId); err != nil {
		respondWithErrorgrp(w, http.StatusInternalServerError, "Errore nell'aggiornamento del gruppo")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "ID gruppo aggiornato con successo"})
}}

// GetMembersHandler restituisce la lista di membri di un gruppo
func GetMembersHandler(database database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId := vars["groupId"]
	members, err := database.GetMembers(groupId)
	if err != nil {
		respondWithErrorgrp(w, http.StatusInternalServerError, "Errore nel recupero dei membri del gruppo")
		return
	}

	respondWithJSON(w, http.StatusOK, members)
}
}

// CountMemberHandler conta il numero di membri in una conversazione di gruppo
func CountMemberHandler(database database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	conversationID := vars["SourceChatId"]

	count, err := database.CountMember(conversationID)
	if err != nil {
		respondWithErrorgrp(w, http.StatusInternalServerError, "Errore nel conteggio dei membri")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]int{"memberCount": count})
}}

// utils.respondWithError invia una risposta JSON di errore
func respondWithErrorgrp(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}


func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}
