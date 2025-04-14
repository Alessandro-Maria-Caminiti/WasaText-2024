package database

// Query for update the group id
var UpdateGroupId  = `UPDATE conversation SET GroupId  = ? WHERE ConversationId = ?`

// Function used to update the group id of a conversation
func (db *appdbimpl) UpdateGroupId (GroupId  string, conversationId string) error {
	_, err := db.c.Exec(UpdateGroupId , GroupId , conversationId)
	return err
}
