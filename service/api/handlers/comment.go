package handlers

import (
	"encoding/json"
	"log" 
	"net/http"
	"wasatext/service/database"
	"wasatext/service/api/structs"
	"github.com/gorilla/mux"
)

// CommentRequest represents the structure of the comment request body
type CommentRequest struct {
	MessageId string `json:"messageId` 
	Content  string `json:"content"`
}

// CommentResponse represents the structure of the response body after commenting
type CommentResponse struct {
	MessageID string `json:"messageId"`
	Status    string `json:"status"`
}

// CommentMessageHandler handles posting a comment on a specific message
func CommentMessageHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		SenderID := vars["id"]
		ConversationID := vars["sourceChatId"]


		var commentReq CommentRequest
		if err := json.NewDecoder(r.Body).Decode(&commentReq); err != nil {
			log.Printf("Error decoding request body: %v", err)
			http.Error(w, `{"error":"Invalid input data"}`, http.StatusBadRequest)
			return
		}

		// Controlli di validazione
		if SenderID == "" || commentReq.MessageId == "" {
			log.Printf("SenderId:%s,Comment lenght:%d,MessageId:%s", SenderID, len(commentReq.Content), commentReq.MessageId)
			http.Error(w, `{"error":"Sender ID, content or MEssageId are required"}`, http.StatusBadRequest)
			return
		}

		// Creazione del commento nel database
		comment := structs.Comment{
			CommentUserId:  SenderID, 
			MessageId:      commentReq.MessageId,           
			Comment:        commentReq.Content,   
			ConversationId:  ConversationID,
		}
		commentID, err := db.CreateComment(comment) 
		if err != nil {
			log.Printf("Error creating comment: %v", err)
			http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(CommentResponse{
			MessageID: commentID.CommentId,
			Status:    "Sent",
		})
	}
}


func DeleteCommentHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		commentID := vars["CommentId"]
		messageID := vars["MessageId"]
		conversationID := r.URL.Query().Get("ConversationID")

		if err := db.DeleteComment(commentID, messageID, conversationID) ; err != nil {
			log.Printf("Error deleting comment: %v", err)
			http.Error(w, `{"error":"Error deleting comment"}`, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func GetMsgCommentHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		messageID := vars["messageId"]
		conversationID := r.URL.Query().Get("ConversationID")
		comments, err := db.GetMsgComments(messageID,conversationID)
		if err != nil {
			log.Printf("Error retrieving comments: %v", err)
			http.Error(w, `{"error":"Error retrieving comments"}`, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(comments)
	}
}

func GetCommentByUserHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID := vars["id"]
		messageID := vars["MessageId"]
		conversationID := r.URL.Query().Get("ConversationID")

		comments, err := db.GetCommentByUser(userID,messageID,conversationID)
		if err != nil {
			log.Printf("Error retrieving user comments: %v", err)
			http.Error(w, `{"error":"Error retrieving user comments"}`, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(comments)
	}
}

func UpdateCommentsHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		commentID := vars["commentId"]
		messageID := vars["MessageId"]
		conversationID := r.URL.Query().Get("ConversationID")

		var updateReq struct {
			Content string `json:"content"`
		}

		if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
			http.Error(w, `{"error":"Invalid input data"}`, http.StatusBadRequest)
			return
		}

		if err := db.UpdateComment(updateReq.Content, commentID, messageID,conversationID ); err != nil {
			log.Printf("Error updating comment: %v", err)
			http.Error(w, `{"error":"Error updating comment"}`, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}