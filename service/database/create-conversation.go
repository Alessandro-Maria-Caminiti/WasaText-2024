package database

import (
	"github.com/google/uuid"
	"wasatext/service/api/structs"
)

// Query used to create a conversation
var AddConversation = `INSERT INTO conversation (ConversationId) VALUES (?);`


func (db *appdbimpl) CreateConversation(c structs.Conversation) (structs.Conversation, error) {
	// New conversation
	var newConv structs.Conversation
	// Set the value of the new conversation
	newUUID := uuid.New().String()
    newConv.ConversationId= newUUID

	// Execute the query to create the conversation
	_, err := db.c.Exec(AddConversation, newConv.ConversationId)
	if err != nil {
		return structs.Conversation{}, err
	}
	// Return the new conversation
	return structs.Conversation{
		ConversationId: newConv.ConversationId,
	}, nil
}
