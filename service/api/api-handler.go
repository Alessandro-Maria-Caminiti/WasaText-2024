package api

import (
	"net/http"
	"log"
	"wasatext/service/api/handlers"
	"github.com/gorilla/mux"
)

// Handler(rt.db) returns an instance of http.Handler(rt.db) that handles API routes
func (rt *_router) Handler() http.Handler {
	// Initialize router
	if rt.db == nil {
        log.Fatal("database is nil in Handler() - this is a programming error")
    }
	router := mux.NewRouter()



	// Session routes
	router.HandleFunc("/session", handlers.LoginHandler(rt.db)).Methods("POST")
	router.HandleFunc("/search/users", handlers.SearchUserHandler(rt.db)).Methods("GET")

	// Chat routes
	router.HandleFunc("/chats/{id}", handlers.RetrieveChatsHandler(rt.db)).Methods("GET")
	router.HandleFunc("/chats/{id}/{sourceChatId}", handlers.GetConversationByIdHandler(rt.db)).Methods("GET")
	router.HandleFunc("/chats/{id}/{sourceChatId}/Agroup", handlers.AddToGroupHandler(rt.db)).Methods("PUT")
	router.HandleFunc("/chats/{id}/{sourceChatId}/leavegroup", handlers.LeaveGroupChatHandler(rt.db)).Methods("DELETE")
	router.HandleFunc("/chats/{id}/{sourceChatId}/Grouphoto", handlers.ChangeGroupPhotoHandler(rt.db)).Methods("PUT")
	router.HandleFunc("/chats/{id}/{sourceChatId}/Groupname", handlers.ChangeGroupNameHandler(rt.db)).Methods("PUT")
	router.HandleFunc("/chats/{id}/creategroup",handlers.CreateGroupHandler(rt.db)).Methods("POST")
	router.HandleFunc("/chats/{id}/{SourceChatId}/Dgroup", handlers.DeleteGroupHandler(rt.db)).Methods("DELETE")
	router.HandleFunc("/chats/{id}/{SourceChatId}/GetGroup", handlers.GetGroupHandler(rt.db)).Methods("GET")
	router.HandleFunc("/chats/{id}/{SourceChatId}/Getconversationgroup", handlers.GetConvGroupHandler(rt.db)).Methods("GET")
	router.HandleFunc("/chats/{id}/{SourceChatId}/UpdateGroupId", handlers.UpdateGroupIdHandler(rt.db)).Methods("PUT")
	router.HandleFunc("/chats/{id}/{SourceChatId}/Members", handlers.GetMembersHandler(rt.db)).Methods("GET")
	router.HandleFunc("/chats/{id}/{SourceChatId}/lastmessage", handlers.UpdateLastMessageHandler(rt.db)).Methods("PUT")
	router.HandleFunc("/chats/{id}/{SourceChatId}/MemberCount", handlers.CountMemberHandler(rt.db)).Methods("GET")
	router.HandleFunc("/chats/{id}/{SourceChatId}/countcheckmarks", handlers.CountCheckmarksHandler(rt.db)).Methods("GET")
	router.HandleFunc("/chats/{id}/CreatePrivateConversation/{destid}", handlers.CreatePrivateConversationHandler(rt.db)).Methods("PUT")
	router.HandleFunc("/chats/{id}/{SourceChatId}/DeletePrivateConversation", handlers.DeletePrivateConversationHandler(rt.db)).Methods("DELETE")
	
	//Messages routes
	router.HandleFunc("/chats/{id}/{sourceChatId}/messages", handlers.SendMessageHandler(rt.db)).Methods("POST")
	router.HandleFunc("/chats/{id}/{sourceChatId}/forward_messages/{targetChatId}", handlers.ForwardMessageHandler(rt.db)).Methods("POST")
	router.HandleFunc("/messages/{SourceChatId}/{messageId}", handlers.GetMessageByIdHandler(rt.db)).Methods("GET")
	router.HandleFunc("/messages/{SourceChatId}", handlers.GetMessagesHandler(rt.db)).Methods("GET")
	router.HandleFunc("/chats/{id}/{sourceChatId}/delete_message/{messageId}", handlers.DeleteMessageHandler(rt.db)).Methods("DELETE")
	router.HandleFunc("/messages/{SourceChatId}/{messageId}/checkmarks", handlers.CountCheckmarksHandler(rt.db)).Methods("GET")
	router.HandleFunc("/messages/{SourceChatId}/{messageId}/update", handlers.UpdateLastMessageHandler(rt.db)).Methods("PUT")
	router.HandleFunc("/messages/{SourceChatId}/{messageId}/insert-checkmark", handlers.InsertCheckmarksHandler(rt.db)).Methods("POST")

	// Comment routes
	router.HandleFunc("/chats/{id}/{sourceChatId}/comment", handlers.CommentMessageHandler(rt.db)).Methods("POST")            // Crea un nuovo commento
	router.HandleFunc("/chats/{id}/{sourceChatId}/comment/{commentId}", handlers.DeleteCommentHandler(rt.db)).Methods("DELETE") // Elimina un commento specifico
	router.HandleFunc("/chats/{id}/{sourceChatId}/comments", handlers.GetMsgCommentHandler(rt.db)).Methods("GET")            // Ottieni tutti i commenti di un messaggio
	router.HandleFunc("/users/{userId}/comments", handlers.GetCommentByUserHandler(rt.db)).Methods("GET")              // Ottieni tutti i commenti di un utente
	router.HandleFunc("/chats/{id}/{sourceChatId}/comment/{commentId}", handlers.UpdateCommentsHandler(rt.db)).Methods("PUT") // Modifica un commento specifico


	// Settings routes
	router.HandleFunc("/settings", handlers.GetUserSettings).Methods("GET")
	router.HandleFunc("/settings", handlers.UpdateUserSettings).Methods("POST")
	router.HandleFunc("/settings/{id}/userphoto", handlers.SetMyPhotoHandler(rt.db)).Methods("PUT")
	router.HandleFunc("/settings/{id}/deleteuser", handlers.DeleteUserHandler(rt.db)).Methods("DELETE")
	router.HandleFunc("/settings/user/{id}/setusername", handlers.SetUsernameHandler(rt.db)).Methods("PUT")


	return router
}
