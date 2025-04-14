package handlers

import (
	"encoding/json"
	"log" 
	"net/http"
	"wasatext/service/database"
	"github.com/gorilla/mux"
)

// GetMessageByIdHandler - Recupera un messaggio specifico in una conversazione
func GetMessageByIdHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		messageIdStr := vars["messageId"]
		convIdStr := vars["convId"]

		message, err := db.GetMessageById(messageIdStr, convIdStr)
		if err != nil {
			log.Printf("Error retrieving message: %v", err)
			http.Error(w, "Error retrieving message", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message)
	}
}

// GetMessagesHandler - Recupera tutti i messaggi di una conversazione
func GetMessagesHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		convIdStr := vars["SourceChatId"]
		log.Printf(convIdStr)

		messages, err := db.GetMessages(convIdStr)
		i := 0
		for (i < len(messages)){

		
		if (messages[i].ConversationId != convIdStr){
			log.Printf("Sono un puzzone")
		}
	i =  i+1}

		if err != nil {
			log.Printf("Error retrieving messages: %v", err)
			http.Error(w, "Error retrieving messages", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(messages)
	}
}

// CountCheckmarksHandler - Conta il numero di checkmark per un messaggio
func CountCheckmarksHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		convIdStr := vars["convId"]
		messageIdStr := vars["messageId"]


		count, err := db.CountCheckmarks(convIdStr, messageIdStr)
		if err != nil {
			log.Printf("Error counting checkmarks: %v", err)
			http.Error(w, "Error counting checkmarks", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]int{"checkmarks": count})
	}
}

// UpdateLastMessageHandler - Aggiorna l'ultimo messaggio di una conversazione
func UpdateLastMessageHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		convIdStr := vars["convId"]
		messageIdStr := vars["messageId"]

		err := db.UpdateLastMessage(messageIdStr, convIdStr)
		if err != nil {
			log.Printf("Error updating last message: %v", err)
			http.Error(w, "Error updating last message", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}
}

// InsertCheckmarksHandler - Inserisce un checkmark per un messaggio
func InsertCheckmarksHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		convIdStr := vars["convId"]
		messageIdStr := vars["messageId"]
		userId := r.URL.Query().Get("userId")

		if userId == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		err := db.InsertCheckmarks(convIdStr, messageIdStr, userId)
		if err != nil {
			log.Printf("Error inserting checkmark: %v", err)
			http.Error(w, "Error inserting checkmark", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "checkmark added"})
	}
}
