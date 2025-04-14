package	database

// Update status of the message
var UpdateStatus = `UPDATE message SET Status = ? WHERE ConversationId = ? AND MessageId = ?`

// UpdateStatus updates the status of the message
func (db *appdbimpl) UpdateStatus(conversationID string, messageID string, status string) error {
	_, err := db.c.Exec(UpdateStatus, status, conversationID, messageID)
	if err != nil {
		return err
	}
	return nil
}