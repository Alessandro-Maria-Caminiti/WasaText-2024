package database

// Query used to update comment
var updatereaction = `UPDATE comment SET Comment = ? WHERE CommentId = ? AND MessageId = ? AND conversationId = ?`

func (db *appdbimpl) UpdateComment(comment string, commentId string, messageId string, conversationId string) error {
	_, err := db.c.Exec(updatereaction, comment, commentId, messageId, conversationId)
	return err
}
