package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"wasatext/service/database"
	"github.com/gorilla/mux"
	"wasatext/service/api/models"
	"wasatext/service/api/images"
	"wasatext/service/api/structs"
)

// User represents a participant in a chat
type User struct {
	userID   string `json:"userId"`
	Username string `json:"username"`
}

// Conversation represents a chat or group conversation
type Conversation struct {
	Identifier           string    `json:"identifier"`
	Username                 string    `json:"name"`
	Participants         []User    `json:"participants"`
	LastMessage          string    `json:"lastMessage"`
	LastMessageTimestamp time.Time `json:"lastMessageTimestamp"`
	IsGroupChat          bool      `json:"isGroupChat"`
}

// Struct used for the response
	type Response struct {
		Conversation structs.Conversation `json:"conversation"`
		User         models.User                 `json:"user"`
		Group        models.Group                `json:"group"`
		GroupUsers   []models.User               `json:"groupUsers"`
		Message      structs.Message      `json:"message"`
		SenderUser   models.User                 `json:"senderUser"`
		User_to_send models.User 		`json:"usertosend"`
	}

// ConversationsResponse represents the response structure for chats	
	func RetrieveChatsHandler(db database.AppDatabase) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			userID := vars["id"]
			if userID == "" {
				http.Error(w, "User ID is required", http.StatusBadRequest)
				return
			}
	
			conversations, err := db.GetUserConversations(userID)
			if err != nil {
				http.Error(w, "Error retrieving chats", http.StatusInternalServerError)
				return
			}
	
			response := make([]Response, len(conversations))
	
			for idx, conv := range conversations {
				if conv.GroupId == "" {
					// Handle individual chat
					
					user, err := db.GetUserById(userID)
					if err != nil {
						http.Error(w, "Error getting user linea 66", http.StatusInternalServerError)
						return
					}
	
					var userModel models.User
					userModel.Username = user.Username
					userModel.UserId = user.UserId
	
					if user.Photo != "" {
						userModel.Photo, _ = images.ImageToBase64(user.Photo)
					} else {
						userModel.Photo, _ = images.ImageToBase64(images.SetDefaultPhoto(user.UserId))
					}
					message, _ := db.GetMessageById(conv.LastMessageId, conv.ConversationId)

					sender, _ := db.GetUserById(message.SenderUserId)
				
					var senderUser models.User
					senderUser.Username = sender.Username
					senderUser.UserId = sender.UserId
	
					if sender.Photo != "" {
						senderUser.Photo, _ = images.ImageToBase64(sender.Photo)
					} else {
						senderUser.Photo, _ = images.ImageToBase64(images.SetDefaultPhoto(sender.UserId))
					}
					users, err := db.SearchUsersfromconv(conv.ConversationId)
					log.Printf("%v", users)
					if err != nil{
						http.Error(w, "Error getting user linea 94", http.StatusInternalServerError)
						return
					}
					var USER models.User
					for fx,userz := range users{
						if (userz.UserId != userID){
							
							USER.Username = userz.Username
							USER.UserId = userz.UserId
							if userz.Photo != "" {
								USER.Photo, _ = images.ImageToBase64(userz.Photo)
							} else {
								USER.Photo, _ = images.ImageToBase64(images.SetDefaultPhoto(userz.UserId))
							}
							log.Printf("%v",fx)
						}
					}
					response[idx] = Response{
						Conversation: conv,
						User:         userModel,
						Message:      message,
						SenderUser:   senderUser,
						User_to_send: USER,
					}
				} else {
					// Handle group chat
					groupDB, err := db.GetGroupById(conv.GroupId)
					if err != nil {
						http.Error(w, "Error retrieving group", http.StatusInternalServerError)
						return
					}
	
					var group models.Group
					group.GroupId = conv.GroupId
					group.GroupName = groupDB.GroupName
	
					if groupDB.GrouPhoto != "" {
						group.Photo, _ = images.ImageToBase64(groupDB.GrouPhoto)
					} else {
						group.Photo, _ = images.ImageToBase64(images.SetDefaultPhoto(conv.GroupId))
					}
					message, _ := db.GetMessageById(conv.LastMessageId, conv.ConversationId)
					sender, _ := db.GetUserById(message.SenderUserId)
	
					var senderUser models.User
					senderUser.Username = sender.Username
					senderUser.UserId = sender.UserId
	
					if sender.Photo != "" {
						senderUser.Photo, _ = images.ImageToBase64(sender.Photo)
					} else {
						senderUser.Photo, _ = images.ImageToBase64(images.SetDefaultPhoto(sender.UserId))
					}
				
					users, err := db.GetMembers(conv.GroupId)
					if err != nil {
						http.Error(w, "Error retrieving group members", http.StatusInternalServerError)
						return
					}
	
					var groupUsers []models.User
					for _, user := range users {
						userDB, err := db.GetUserById(user.UserId)
						if err != nil {
							http.Error(w, "Error retrieving user linea 158", http.StatusInternalServerError)
							return
						}
	
						var userModel models.User
						userModel.Username = userDB.Username
						userModel.UserId = userDB.UserId
	
						if userDB.Photo != "" {
							userModel.Photo, _ = images.ImageToBase64(userDB.Photo)
						} else {
							userModel.Photo, _ = images.ImageToBase64(images.SetDefaultPhoto(userDB.UserId))
						}
	
						groupUsers = append(groupUsers, userModel)
					}
					var userp models.User
					response[idx] = Response{
						Conversation: conv,
						Group:        group,
						GroupUsers:   groupUsers,
						Message:      message,
						SenderUser:   senderUser,
						User_to_send: userp,
					}
				}
			}
	
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, "Error encoding response", http.StatusInternalServerError)
			}
		}
	}
	

// GetConversationByIdHandler - Recupera una conversazione specifica
func GetConversationByIdHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		convIdStr := vars["sourceChatId"]
		log.Printf(convIdStr)
		conversation, err := db.GetConversationById(convIdStr)
		if err != nil {
			log.Printf("Error retrieving conversation: %v", err)
			http.Error(w, "Error retrieving conversation", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(conversation)
	}
}
