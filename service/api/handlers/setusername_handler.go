package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"wasatext/service/api/models"
	"wasatext/service/database"
	"strings"
)

// SetUsernameHandler gestisce la modifica dell'username, mantenendo lo stesso identifier.
func SetUsernameHandler(db database.AppDatabase) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Estrai l'userID dall'URL
           // Per la route settings/{id}/deleteuser
   		 parts := strings.Split(r.URL.Path, "/")
    
		// Controlliamo che ci siano abbastanza parti nell'URL e che l'ID non sia vuoto
		// L'URL dovrebbe essere del tipo /settings/{id}/deleteuser
		// quindi l'ID si trova nella posizione parts[2]
		fmt.Printf(parts[3])
		if len(parts) < 4 || parts[3] == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Missing user ID"})
			return
		}
		userID := parts[3]

        if userID == "" {
            http.Error(w, "User ID is required", http.StatusBadRequest)
            return
        }

        var user models.UserRequest
        // Decode il body JSON
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, "Invalid JSON format", http.StatusBadRequest)
            return
        }

        // Validiamo il nuovo username
        if err := validateUsername(user.Name); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Controlla se l'utente esiste
        _, err := db.CheckIfExist(user.Name)  // Assumendo che hai questa funzione
        if err != nil {
            http.Error(w, "User already exists", http.StatusNotFound)
            return
        }

        // Aggiorniamo il nome utente
        if err := db.SetMyUsername(userID, user.Name); err != nil {
            http.Error(w, "Error updating username", http.StatusInternalServerError)
            return
        }
		 userdb,err := db.GetUserByName(user.Name)
		 if err != nil{
			http.Error(w, "User not found", http.StatusNotFound)
            return
		 }
        // Rispondiamo con l'identifier che rimane invariato
        response := map[string]string{"identifier": userID, "Name":userdb.Username}
        w.WriteHeader(http.StatusOK)  // Meglio usare StatusOK per un aggiornamento riuscito
        json.NewEncoder(w).Encode(response)
    }
}
// validateUsername validates the username against the given pattern and length
func validateUsername(username string) error {
	// Check length
	if len(username) < 3 || len(username) > 32 {
		return fmt.Errorf("Username must be between 3 and 32 characters long.")
	}

	// Check for valid characters using regular expression
	// This regex allows only alphanumeric characters, dots, and underscores.
	var validUsernamePattern = `^[a-zA-Z0-9._]+$`
	matched, _ := regexp.MatchString(validUsernamePattern, username)
	if !matched {
		return fmt.Errorf("Username contains invalid characters or does not meet length requirements.")
	}

	return nil
}
