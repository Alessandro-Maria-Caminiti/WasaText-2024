package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "wasatext/service/database"
    "github.com/google/uuid"
    "github.com/gorilla/mux"
)

// AddToPrivateChatRequest rappresenta il corpo della richiesta previsto
type AddToPrivateChatRequest struct {
    UserID string `json:"UserId"`
}

// AddToPrivateChatResponse rappresenta una risposta di successo
type AddToPrivateChatResponse struct {
    Status string `json:"status"`
}

// AddToPrivateChatHandler gestisce l'aggiunta di un utente a una chat privata
func AddToPrivateChatHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    sourceChatID := vars["SourceChatId"]
    
    // Analizza e convalida il corpo della richiesta
    var addReq AddToPrivateChatRequest
    if err := json.NewDecoder(r.Body).Decode(&addReq); err != nil {
        log.Printf("Errore nella decodifica del corpo della richiesta: %v", err)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Dati di input non validi"})
        return
    }
    
    // Convalida userId
    if addReq.UserID == "" {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "UserID è obbligatorio"})
        return
    }

    // Convalida formato UUID
    _, err := uuid.Parse(addReq.UserID)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Formato ID utente non valido"})
        return
    }

    _, err = uuid.Parse(sourceChatID) // Usa sourceChatID invece di ConversationId
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Formato ID chat non valido"})
        return
    }

    // Aggiungi utente alla chat privata utilizzando la funzione del database
    err = db.AddUserConv(sourceChatID,addReq.UserID) // Usa addReq.UserID invece di UserId
    if err != nil {
        log.Printf("Errore database: %v", err)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "Impossibile aggiungere l'utente alla chat privata"})
        return
    }
    // Invia risposta di successo
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(AddToPrivateChatResponse{
        Status: "Utente aggiunto alla chat privata con successo",
    })
}}