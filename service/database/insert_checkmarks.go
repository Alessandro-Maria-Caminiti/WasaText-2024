package database

var insertcheckmarks = `INSERT INTO checkmarks (ConversationId, MessageId, UserId) VALUES (?, ?, ?)`

// InsertCheckmarks inserts checkmarks into the database
func (db *appdbimpl) InsertCheckmarks(conversationID string, messageID string, UserId string) error {
	_, err := db.c.Exec(insertcheckmarks, conversationID, messageID, UserId)
	if err != nil {
		return err
	}
	return nil
}