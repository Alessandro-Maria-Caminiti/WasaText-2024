package database

// Query used to add in the conversation_user table all the user in the conversation table
var AddUserConv = `INSERT INTO conversation_user (ConversationId, UserId) VALUES (?, ?)`

func (db *appdbimpl) AddUserConv(conversationId string, UserId string) error {
	_, err := db.c.Exec(AddUserConv, conversationId, UserId)
	if err != nil {
		return err
	}
	return nil
}
