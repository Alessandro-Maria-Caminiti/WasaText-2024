package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"wasatext/service/database"
)

// SearchHandler gestisce la ricerca di utenti nel database
func SearchUserHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("name") // Ottiene il parametro di ricerca dalla query string

		if query == "" {
			http.Error(w, "Search query is required", http.StatusBadRequest)
			return
		}

		// Chiama la funzione SearchUsers dal database
		users, err := db.SearchUsers(query)
		if err != nil {
			log.Printf("Error searching users: %v", err)
			http.Error(w, "Error searching users", http.StatusInternalServerError)
			return
		}
		
		// Invia la risposta JSON con i risultati della ricerca
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}
