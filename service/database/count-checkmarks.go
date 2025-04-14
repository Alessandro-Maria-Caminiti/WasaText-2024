package database

// Query used to count member of a checkmarks
var CountCheckmarks = `SELECT COUNT(UserId) FROM checkmarks WHERE ConversationId = ? AND MessageId = ?`

// CountCheckmarks returns the number of member of a checkmarks
func (db *appdbimpl) CountCheckmarks(conversationID string, messageID string) (int, error) {
	var count int
	err := db.c.QueryRow(CountCheckmarks, conversationID, messageID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}