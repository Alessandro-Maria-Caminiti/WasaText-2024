package database

// Query used for search a conversation of a group
var GetConversationGroup = `SELECT GroupId FROM group_t WHERE ConversationId  = ?`

// Function used to get a conversation of a group
func (db *appdbimpl) GetConvGroup(ConversationId  string) (string, error) {
	var GroupId string
	err := db.c.QueryRow(GetConversationGroup, ConversationId).Scan(&GroupId)
	if err != nil {
		return "", err
	}
	return GroupId, nil
}
