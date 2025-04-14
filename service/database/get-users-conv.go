package database

// Query used to get user of a conversation
var getusersconversetion = `SELECT UserId FROM conversation_user WHERE ConversationId = ? AND UserId != ?`

// Function
func (db *appdbimpl) GetUsersConv(ConversationId string, UserId string) (User, error) {
	var user User
	err := db.c.QueryRow(getusersconversetion, ConversationId, UserId).Scan(&user.UserId)
	return user, err
}
