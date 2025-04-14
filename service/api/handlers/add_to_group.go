package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "wasatext/service/api/models"
    "wasatext/service/database"
    "github.com/gorilla/mux"
    "github.com/google/uuid"
)

// AddToGroupRequest rappresenta la richiesta che contiene una lista di utenti
type AddToGroupRequest struct {
    Users []models.User `json:"users"` // Sempre un array, anche per un solo utente
}

// AddToGroupResponse rappresenta la risposta con i risultati
type AddToGroupResponse struct {
    Status  string   `json:"status"`
    AddedUsers []models.User `json:"members"` 
    Errors []string `json:Errors`

}

// AddToGroupHandler gestisce l'aggiunta di utenti a un gruppo
func AddToGroupHandler(db database.AppDatabase) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        sourceChatId := vars["sourceChatId"]
        
        // Validare il formato UUID del gruppo
        if _, err := uuid.Parse(sourceChatId); err != nil {
            
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "Invalid group ID format"})
            return
        }
        log.Printf(sourceChatId)
        groupId, err := db.GetConvGroup(sourceChatId)
        log.Printf("%v", groupId)
        if err != nil{
            log.Printf("Error finding conv: %v", err)
		http.Error(w, "Error finding conv", http.StatusInternalServerError)
		return
        }
        // Decodificare il corpo della richiesta
        var addReq AddToGroupRequest
        if err := json.NewDecoder(r.Body).Decode(&addReq); err != nil {
            log.Printf("Error decoding request body: %v", err)
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "Invalid input data"})
            return
        }
        // Controllare se ci sono utenti nella richiesta
        if len(addReq.Users) == 0 {
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "At least one user is required"})
            return
        }

        var addedUsers []models.User
        var errors []string

        // Iterare su ogni utente e aggiungerlo al gruppo
        for _, user := range addReq.Users {
            if user.Username == "" {
                errors = append(errors, "Username is required for user ID "+user.UserId)
                continue
            }

            if _, err := uuid.Parse(user.UserId); err != nil {
                errors = append(errors, "Invalid User ID format for user "+user.UserId)
                continue
            }
            

            err = db.AddUserGroup(user.UserId, groupId)
            if err != nil {
                log.Printf("Database error for user %s: %v", user.UserId, err)
                errors = append(errors, "Failed to add user "+user.UserId+" to group")
                continue
            }

            addedUsers = append(addedUsers, user)
        }

        // Preparare la risposta
        response := AddToGroupResponse{
            Status:  "Processing completed",
            AddedUsers: addedUsers,
            Errors: errors,
        }

        w.Header().Set("Content-Type", "application/json")
        if len(errors) > 0 {
            w.WriteHeader(http.StatusMultiStatus)
        } else {
            w.WriteHeader(http.StatusCreated)
        }
        json.NewEncoder(w).Encode(response)
    }
}
