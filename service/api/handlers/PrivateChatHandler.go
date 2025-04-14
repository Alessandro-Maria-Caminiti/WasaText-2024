package handlers

import (
	"encoding/json"
	"net/http"
	"wasatext/service/database"
	"wasatext/service/api/structs"
	"github.com/gorilla/mux" 
) 

// CreatePrivateConversationHandler gestisce la creazione di una conversazione privata
func CreatePrivateConversationHandler(database database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		UserId := vars["id"]
		destid := vars["destid"]

		var conversation structs.Conversation
		newConv, err := database.CreateConversation(conversation)
		if err != nil {
			http.Error(w, "Errore nella creazione della conversazione", http.StatusInternalServerError)
			return
		}

		err = database.AddUserConv(newConv.ConversationId, UserId)
		if err != nil {
			http.Error(w, "Errore nell'aggiunta dell'utente alla conversazione", http.StatusInternalServerError)
			return
		}


		err = database.AddUserConv(newConv.ConversationId, destid)
		if err != nil {
			http.Error(w, "Errore nell'aggiunta dell'utente alla conversazione", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		// Ecco il fix
		json.NewEncoder(w).Encode(map[string]interface{}{
			"conversationId": newConv.ConversationId,
		})
	}
}


// DeletePrivateConversationHandler elimina una conversazione privata
func DeletePrivateConversationHandler(database database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	convId:= vars["conversationID"]

	err := database.DeleteConv(convId)
	if err != nil {
		http.Error(w, "Errore nell'eliminazione della conversazione", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
}