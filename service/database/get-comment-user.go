package database

import (
	"wasatext/service/api/structs"
)

// Query used to take the comment of a user
var Getreactionbyuser = `SELECT CommentId, Comment FROM comment WHERE CommentUserId = ? AND MessageId = ? AND ConversationId = ?`

func (db *appdbimpl) GetCommentByUser(UserId string, messageId string, ConversationId string) (structs.Comment, error) {
	var comment structs.Comment
	err := db.c.QueryRow(Getreactionbyuser, UserId, messageId, ConversationId).Scan(&comment.CommentId, &comment.Comment)
	return comment, err
}
