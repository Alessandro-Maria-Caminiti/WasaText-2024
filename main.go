package main

import (
	"log"
	"net/http"

	"wasatext/handlers"
	"wasatext/routes" // Import routes package

	"github.com/gorilla/mux"
)

func main() {
	// Initialize router
	router := mux.NewRouter()

	// Load routes
	routes.RegisterSessionRoutes(router) // Session routes
	routes.RegisterUserRoutes(router)    // User routes
	router.HandleFunc("/session", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/settings/setusername", handlers.SetUsernameHandler).Methods("POST")
	router.HandleFunc("/chats", handlers.RetrieveChatsHandler).Methods("GET")
	router.HandleFunc("/chats/{sourceChatId}", handlers.RetrieveChatDetailsHandler).Methods("GET")
	router.HandleFunc("/chats/{sourceChatId}/messages", handlers.SendMessageHandler).Methods("POST")
	router.HandleFunc("/chats/{sourceChatId}/forward_messages/{targetChatId}", handlers.ForwardMessageHandler)
	router.HandleFunc("/chats/{sourceChatId}/delete_message/{messageId}", handlers.DeleteMessageHandler)
	router.HandleFunc("/chats/{sourceChatId}/Agroup", handlers.AddToGroupHandler).Methods("POST")
	router.HandleFunc("/chats/{sourceChatId}/comment", handlers.CommentMessageHandler).Methods("POST")
	router.HandleFunc("/chats/{sourceChatId}/uncomment/{commentId}", handlers.DeleteCommentHandler).Methods("DELETE")
	router.HandleFunc("/chats/{sourceChatId}/Dgroup", handlers.LeaveGroupChatHandler).Methods("DELETE")
	router.HandleFunc("/chats/{sourceChatId}/Grouphoto", handlers.ChangeGroupPhotoHandler).Methods("POST")
	router.HandleFunc("/chats/{sourceChatId}/Groupname", handlers.ChangeGroupNameHandler).Methods("POST")
	router.HandleFunc("/settings", handlers.GetUserSettings).Methods("GET")
	router.HandleFunc("/settings", handlers.UpdateUserSettings).Methods("POST")
	router.HandleFunc("/settings/userphoto", handlers.SetUserPhoto).Methods("POST")
	// Start the server
	log.Println("Server is listening on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
