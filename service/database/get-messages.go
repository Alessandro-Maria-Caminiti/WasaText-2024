package database

import (
	"wasatext/service/api/structs"
	"log"
)

// Query used to get the messages of a conversation from the database
var getmessageid = `SELECT ConversationId,MessageId, Message, Status, UserId, SendTime, COALESCE(Photo, "") FROM message WHERE ConversationId = ?`

func (db *appdbimpl) GetMessages(ConversationId string) ([]structs.Message, error) {
	// Create a new slice of messages
	var messages []structs.Message
	// Execute the query to get the messages of a conversation
	rows, err := db.c.Query(getmessageid, ConversationId)
	if err != nil {

		return nil, err
	}
	defer func() { rows.Close() }()

	// Iterate
	for rows.Next() {
		if rows.Err() != nil {
			return nil, err
		}
		// Create a new message
		var message structs.Message
		// Scan the values of the message
		err := rows.Scan(&message.ConversationId,&message.MessageId, &message.Text, &message.Status, &message.SenderUserId, &message.SendTime, &message.Photo)
		if err != nil {
			log.Printf(err.Error())
			return nil, err
		}
		// Append the message to the slice
		messages = append(messages, message)
	}
	return messages, nil
}
