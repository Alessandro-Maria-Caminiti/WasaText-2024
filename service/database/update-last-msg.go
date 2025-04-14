package database

import (
	"database/sql"
	"strconv"
)

// Query used to update the last message from a conversation
var updatelastmessage = `UPDATE conversation SET LastMessageId = ? WHERE ConversationId = ?`


// UpdateLastMessage aggiorna l'ultimo messaggio di una conversazione
func (db *appdbimpl) UpdateLastMessage(messageId string, conversationId string) error {
	// Converti messageId in int64
	msgId, err := strconv.ParseInt(messageId, 10, 64)
	if err != nil {
		return err // Ritorna l'errore se la conversione fallisce
	}

	// Definisci il valore per il database
	value := sql.NullInt64{Int64: msgId, Valid: messageId != ""}

	// Esegui la query di aggiornamento
	_, err = db.c.Exec(updatelastmessage, value, conversationId)
	return err
}