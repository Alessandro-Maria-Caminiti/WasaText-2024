package database

// Query used to count member of a conversation
var CountMember = `SELECT COUNT(UserId) FROM conversation_user WHERE ConversationId = ?`

// CountMember returns the number of member of a conversation
func (db *appdbimpl) CountMember(conversationID string) (int, error) {
	var count int
	err := db.c.QueryRow(CountMember, conversationID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}