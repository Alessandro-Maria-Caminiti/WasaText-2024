package database

import (
	"wasatext/service/api/structs"
)

// Query used to find a message by its id in the database
var GetMessageByid = `SELECT MessageId, Message, SenderUserId, SendTime, ConversationId, COALESCE(Photo, "") FROM message WHERE MessageId = ? AND ConversationId = ?`

func (db *appdbimpl) GetMessageById(messageId string, ConversationId string) (structs.Message, error) {
	var message structs.Message
	err := db.c.QueryRow(GetMessageByid, messageId, ConversationId).Scan(&message.MessageId, &message.Text, &message.SenderUserId, &message.SendTime, &message.ConversationId, &message.Photo)
	return message, err
}
