package database
import (
    "wasatext/service/api/structs"
    "github.com/google/uuid"
)

// Query used to insert new comment in the db
var AddComment = `INSERT INTO comment (CommentId, UserId, Comment, MessageId, ConversationId) VALUES (?, ?, ?, ?, ?)`

func (db *appdbimpl) CreateComment(c structs.Comment) (structs.Comment, error) {
    // New comment
    var newComment structs.Comment
    
    // Generate unique comment ID
    newUUID := uuid.New().String()
    
    // Set values of new comment
    newComment.CommentId = newUUID
    newComment.CommentUserId = c.CommentUserId
    newComment.Comment = c.Comment
    newComment.MessageId = c.MessageId
    newComment.ConversationId = c.ConversationId  // Usa l'ID conversazione ricevuto, non ne generare uno nuovo
    
    // Exec query
    _, err := db.c.Exec(AddComment, newComment.CommentId, newComment.CommentUserId, newComment.Comment, newComment.MessageId, newComment.ConversationId)
    if err != nil {
        return structs.Comment{}, err
    }
    
    return newComment, nil
}