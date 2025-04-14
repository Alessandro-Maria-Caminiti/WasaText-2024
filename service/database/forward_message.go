package database
import (
    "errors"
    "github.com/google/uuid"
    "wasatext/service/api/structs"
    "log"
    "time"
)

// ForwardMessage copia un messaggio esistente in una nuova conversazione
func (db *appdbimpl) ForwardMessage(messageId string, sourceChatId string, targetChatId string, senderId string) (structs.Message, error) {
    // Recupera il messaggio originale
    var originalMessage structs.Message
    query := "SELECT MessageId, Message, Status, UserId, Photo, SendTime FROM message WHERE MessageId = ? AND ConversationId = ?"
    err := db.c.QueryRow(query,
        messageId, sourceChatId).Scan(
        &originalMessage.MessageId,
        &originalMessage.Text,
        &originalMessage.Status,
        &originalMessage.SenderUserId,
        &originalMessage.Photo,
        &originalMessage.SendTime)
    
    if err != nil {
        log.Printf("Errore nel recupero del messaggio con ID %s: %v", messageId, err)
        return structs.Message{}, errors.New("messaggio originale non trovato")
    }
    
    // Genera un nuovo ID per il messaggio inoltrato
    newMessageId := uuid.New().String()
    
    // Ottieni il tempo corrente per l'inoltro
    currentTime := time.Now()
    
    // Crea il nuovo messaggio per il targetChatId
    newMessage := structs.Message{
        MessageId:      newMessageId,
        Text:           originalMessage.Text,
        SendTime:       currentTime,
        Status:         "forwarded",
        SenderUserId:   senderId,
        ConversationId: targetChatId,
        Photo:          originalMessage.Photo,
    }
    
    // Salva il nuovo messaggio nel database
    _, err = db.c.Exec(
        "INSERT INTO message (MessageId, Message, SendTime, Status, UserId, ConversationId, Photo) VALUES (?, ?, ?, ?, ?, ?, ?)",
        newMessage.MessageId,
        newMessage.Text,
        newMessage.SendTime,
        newMessage.Status,
        newMessage.SenderUserId,    // Assicurati che corrisponda al campo UserId del database
        newMessage.ConversationId,  // Assicurati che questo sia nella posizione corretta
        newMessage.Photo)
    
    if err != nil {
        log.Printf("Errore nel messaggio inoltrato - SenderUserId: %s, ConversationId: %s",
            newMessage.SenderUserId, newMessage.ConversationId)
        log.Printf("Errore nell'inserimento del messaggio inoltrato: %v", err)
        return structs.Message{}, err
    }
    
    return newMessage, nil
}