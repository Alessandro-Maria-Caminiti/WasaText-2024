package database

import (
	"github.com/google/uuid"
	"wasatext/service/api/structs"
	"log"
)

// Query used to add the message in the database
var AddMessage = `INSERT INTO message (MessageId, Message, Status, ConversationId, UserId, Photo) VALUES (?, ?, ?, ?, ?, ?)`

func (db *appdbimpl) CreateMessage(msg structs.Message) (structs.Message, error) {
	// New message
	var newMsg structs.Message
	
	// Set the value of the new message
	newMsg.Text = msg.Text
	newUUID := uuid.New().String()
	newMsg.MessageId = newUUID
	newMsg.Status = msg.Status
	newMsg.SenderUserId = msg.SenderUserId
	newMsg.Photo = msg.Photo
	newMsg.ConversationId = msg.ConversationId
	log.Printf("Text %s", newMsg.Text, "MessageId%s", newMsg.MessageId, "Status%s", newMsg.Status, "SenderUSerId%s", newMsg.SenderUserId,"Photo%s", newMsg.Photo, "ConversationId%s", newMsg.ConversationId )
	// Execute the query to create the conversation
	_, err := db.c.Exec(AddMessage, newMsg.MessageId, newMsg.Text, newMsg.Status, newMsg.ConversationId, newMsg.SenderUserId, newMsg.Photo)
	
	if err != nil {
		log.Printf(err.Error())
		return structs.Message{}, err
	}

	// Return the new conversation
	return structs.Message{
		MessageId:      newMsg.MessageId,
		Text:           newMsg.Text,
		SendTime:       newMsg.SendTime,
		Status:         newMsg.Status,
		SenderUserId:   newMsg.SenderUserId,
		ConversationId: newMsg.ConversationId,
		Photo:          newMsg.Photo,
	}, nil
}
