package database

import (
	"wasatext/service/api/structs"
	"log"
)

// Query used to take the conversation from the db with the id
var GetConversationbyId = `SELECT ConversationId FROM conversation WHERE ConversationId = ?`

func (db *appdbimpl) GetConversationById(ConversationId string) (structs.Conversation, error) {
	var conv structs.Conversation
	// Exec the query
	err := db.c.QueryRow(GetConversationbyId, ConversationId).Scan(&conv.ConversationId)
	if err != nil {
		return structs.Conversation{}, err
	}
	groupid, err := db.GetConvGroup(ConversationId)
        if err != nil{
            log.Printf("Errore in GetConvGroup per ID %s: %v", ConversationId, err)
               
        }
        if (groupid != ""){
            conv.GroupId = groupid
        }
        log.Printf("Convid: %s, Conv.Groupid: %s, Conv.Lastmessageid: %s", conv.ConversationId, conv.GroupId, conv.LastMessageId)
        // Stampare i dettagli della conversazione
	log.Printf(conv.GroupId)
	return conv, nil
}
 