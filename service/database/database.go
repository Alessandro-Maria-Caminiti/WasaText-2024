/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"wasatext/service/api/structs"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// -- USER OPERATION -- //

	// Creation of new user in the user table
	CreateUser(u User) (User, error)

	DeleteUser(UserId string) error

	UpdateUserProfileImage(userID string, photoURL string) error
	
 	// Change the username of the user
	SetMyUsername(UserId string, newUsername string) error

	// Check if the username is alredy used
	CheckIfExist(username string) (bool, error)

	// Get User information from the db with the username
	GetUserByName(username string) (User, error)

	// Get User information from the db with the id
	GetUserById(UserId string) (User, error)

	// Delete a member from a group
	LeaveGroup(UserId string, GroupId  string) error

	// Get all members of a group
	GetMembers(GroupId  string) ([]User, error)

	// Set new group name
	SetGroupName(GroupId  string, newName string) error

	// Search
	SearchUsers(search string) ([]User, error)

	// -- GROUP OPERATION -- //

	// Add a user to a group
	AddUserGroup(UserId string, GroupId  string) error

	// Set new group photo
	SetGroupPhoto(GroupId  string, photoUrl string) error

	// Create Group
	CreateGroup(GroupName string, UserId string) (Group, error)

	// Get Groiup information from the db with the id
	GetGroupById(GroupId  string) (Group, error)

	// Check if a user is member of a group
	CheckMember(UserId string, GroupId  string) (bool, error)

	// Delete group
	DeleteGroup(GroupId  string) error

	// Delete all the user from the user_group table there are member of the group
	DeleteMember(GroupId  string, tx *sql.Tx) error
	
	// Checks if the group Exists
	GroupExists(GroupId  string) (bool, error)
	
	DeleteUserFromGroup(UserId string, GroupId  string) error

	// -- CONVERSATION OPERATION -- //


	SearchUsersfromconv(convid string) ([]User, error)
	// Create new Conversation
	CreateConversation(c structs.Conversation) (structs.Conversation, error)

	// Update last message of a conversation
	UpdateLastMessage(messageId string, conversationId string) error

	// Add user and conversation link in the conversation_user table
	AddUserConv(conversationId string, UserId string) error

	// Check if exist a conversation between two users
	CheckIfExistConv(sender string, receiver string) (bool, error)

	// Get the conversation from the id
	GetConversationById(ConversationId string) (structs.Conversation, error)

	// Check if the user is part of the conversation
	CheckUserConv(UserId string, ConversationId string) (bool, error)

	// Get all conversations of a user
	GetUserConversations(UserId string) ([]structs.Conversation, error)

	// Update the group id of a conversation
	UpdateGroupId(GroupId  string, conversationId string) error

	// Get conversation by group id
	GetConvGroup(GroupId  string) (string, error)

	// Remove user from conversation
	DeleteUserConv(UserId string, ConversationId string) error

	// Get users of a conversation
	GetUsersConv(ConversationId string, UserId string) (User, error)

	// Delete conversation
	DeleteConv(ConversationId string) error

	// Get conversation id from sender and receiver
	GetConversation(sender string, receiver string) (string, error)

	// Count the number of member of a conversation
	CountMember(conversationID string) (int, error)

	UpdateStatus(conversationID string, messageId string, status string) error

	// -- MESSAGE OPERATION -- //

	// Create new message
	CreateMessage(msg structs.Message) (structs.Message, error)

	// Get message from the id of the message and the conversation id
	GetMessageById(messageId string, ConversationId string) (structs.Message, error)

	// Get all messages of a conversation
	GetMessages(ConversationId string) ([]structs.Message, error)


	// Delete the message
	DeleteMessage(messageId string, ConversationId string) error

	ForwardMessage(messageId string, sourceChatId string, targetChatId string, senderId string) (structs.Message, error)

	// -- COMMENT OPERATION -- //

	// Create new comment
	CreateComment(c structs.Comment) (structs.Comment, error)

	// Get comment by user id
	GetCommentByUser(UserId string, messageId string, ConversationId string) (structs.Comment, error)

	// Update comment
	UpdateComment(comment string, commentId string, messageId string, conversationId string) error

	// Delete comment
	DeleteComment(commentId string, messageId string, ConversationId string) error

	// Get all comments of a message
	GetMsgComments(msgId string, ConversationId string) ([]structs.RspComment, error)

	// -- CHECKMARKS OPERATION -- //
	InsertCheckmarks(conversationID string, messageId string, UserId string) error

	CountCheckmarks(conversationID string, messageId string) (int, error)


	Ping() error
}

type appdbimpl struct {
	c   *sql.DB
	ctx context.Context
}


// GetInstance returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func GetInstance(db *sql.DB) (AppDatabase, error) {

	// Check if the database is nil (required)
	if db == nil {
		return nil, errors.New("database connection is required")
	}

	// Check if the database is empty
	var tableCount int
	err := db.QueryRow("SELECT COUNT(name) FROM sqlite_master WHERE type='table'").Scan(&tableCount)
	if err != nil {
		return nil, fmt.Errorf("error checking database tables: %w", err)
	}

	// Expected number of tables
	const expectedTables = 8 // Cambia questo numero se necessario

	if tableCount != expectedTables {
		tables := map[string]string{
			"user":               userTableSQL,
			"message":            messageTableSQL,
			"group":              groupTableSQL,
			"user_group":         userGroupTableSQL,
			"conversation":       conversationTableSQL,
			"conversation_user":  conversationUsersSQL,
			"comment":            commentTableSQL,
			"checkmarks":         checkmarksTableSQL,
		}

		// Creazione delle tabelle mancanti
		for name, sqlStatement := range tables {
			_, err = db.Exec(sqlStatement)
			if err != nil {
				return nil, fmt.Errorf("error creating table %s: %w", name, err)
			}
		}
	}

	return &appdbimpl{
		c:   db,
		ctx: context.Background(),
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
