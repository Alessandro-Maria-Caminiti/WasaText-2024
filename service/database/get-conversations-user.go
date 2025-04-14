package database
import (
    "wasatext/service/api/structs"
    "log"
    "encoding/json"
)

// Query used to get all conversations of a user
var GetConversation = `SELECT ConversationId FROM conversation_user WHERE UserId = ?`

func (db *appdbimpl) GetUserConversations(UserId string) ([]structs.Conversation, error) {
    // Log the user ID we're searching for
    log.Printf("Cercando conversazioni per l'utente: %s", UserId)
    
    // Exec query
    rows, err := db.c.Query(GetConversation, UserId)
    if err != nil {
        log.Printf("Errore nella query: %v", err)
        return nil, err
    }
    defer func() { _ = rows.Close() }() 
    
    // Raccogliere tutti gli ID prima
    var convIDs []string
    for rows.Next() {
        var convID string
        err = rows.Scan(&convID)
        if err != nil {
            log.Printf("Errore di scan: %v", err)
            return nil, err
        }
        convIDs = append(convIDs, convID)
    }
    
    if err = rows.Err(); err != nil {
        log.Printf("Errore nelle righe: %v", err)
        return nil, err
    }
    
  
    
    // All conversations
    var convs []structs.Conversation
    
    // Fetch each conversation
    for _, convID := range convIDs {
        conv, err := db.GetConversationById(convID)
        if err != nil {
            log.Printf("Errore in GetConversationById per ID %s: %v", convID, err)
            // Opzionale: invece di restituire errore, potremmo continuare con le altre conversazioni
            // return nil, err
            continue
        }
        groupid, err := db.GetConvGroup(convID)
        if err != nil{
            log.Printf("Errore in GetConvGroup per ID %s: %v", convID, err)
               
        }
        if (groupid != ""){
            conv.GroupId = groupid
        }
        log.Printf("Convid: %s, Conv.Groupid: %s, Conv.Lastmessageid: %s", conv.ConversationId, conv.GroupId, conv.LastMessageId)
        // Stampare i dettagli della conversazione
        convJSON, _ := json.MarshalIndent(conv, conv.GroupId, "  ")
        log.Printf("Dettagli conversazione %s: %s", convID, string(convJSON))
        if conv.ConversationId != "" {
            convs = append(convs, conv)
        }        
        log.Printf("Aggiunta conversazione ID: %s (totale finora: %d)", conv.ConversationId, len(convs))
    }
    
    // Stampa riassuntiva
    log.Printf("Totale conversazioni recuperate: %d", len(convs))
    
    return convs, nil
}