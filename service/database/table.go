package database

// In this file we define the SQL for the tables we want to create in the database.

/*
* USER TABLE SQL
*/
var userTableSQL = `CREATE TABLE IF NOT EXISTS user (
	UserId STRING NOT NULL UNIQUE,
	Username STRING NOT NULL UNIQUE,
  	profile_image STRING,
	PRIMARY KEY(UserId)
);`

/*
* MESSAGE TABLE SQL
*/
var messageTableSQL = `CREATE TABLE IF NOT EXISTS message (
	MessageId STRING NOT NULL,
  	Message STRING,
  	SendTime DATETIME DEFAULT CURRENT_TIMESTAMP,
  	Status TEXT,
  	ConversationId STRING NOT NULL,
  	UserId STRING NOT NULL,
  	Photo STRING,
  	PRIMARY KEY(MessageId, ConversationId),
  	CONSTRAINT fk_message
    	FOREIGN KEY (UserId) REFERENCES user(UserId)
      	ON DELETE CASCADE,
    	FOREIGN KEY (ConversationId) REFERENCES conversation(ConversationId)
      	ON DELETE CASCADE
);`

/*
* GROUP TABLE SQL
*/
var groupTableSQL = `CREATE TABLE IF NOT EXISTS group_t (
ConversationId STRING NOT NULL,	
GroupId STRING NOT NULL,
	GroupName STRING NOT NULL,
	GrouPhoto STRING NOT NULL,
	PRIMARY KEY(GroupId)
	CONSTRAINT fk_group_t
    	FOREIGN KEY (ConversationId) REFERENCES conversation(ConversationId)
      	ON DELETE CASCADE
);`

var userGroupTableSQL = `CREATE TABLE IF NOT EXISTS user_group (
	GroupId STRING NOT NULL,
	UserId STRING NOT NULL,
  	PRIMARY KEY(GroupId, UserId),
	CONSTRAINT fk_user_group
		FOREIGN KEY (GroupId) REFERENCES group_t(GroupId)
		  ON DELETE CASCADE,
		FOREIGN KEY (UserId) REFERENCES user(UserId)
		  ON DELETE CASCADE
);`

/*
* CONVERSATION TABLE SQL
*/
var conversationTableSQL = `CREATE TABLE IF NOT EXISTS conversation (
	ConversationId STRING NOT NULL,
  	GroupId STRING,
  	PRIMARY KEY(ConversationId),
  	CONSTRAINT fk_conversation
    	FOREIGN KEY (GroupId) REFERENCES group_t(GroupId)
      	ON DELETE CASCADE
);`

var conversationUsersSQL = `CREATE TABLE IF NOT EXISTS conversation_user (
  	ConversationId STRING NOT NULL,
  	UserId STRING NOT NULL,
  	PRIMARY KEY(ConversationId, UserId),
  	CONSTRAINT fk_conversation_user
    	FOREIGN KEY (ConversationId) REFERENCES conversation(ConversationId)
      	ON DELETE CASCADE,
    	FOREIGN KEY (UserId) REFERENCES user(UserId)
      	ON DELETE CASCADE
);`

var checkmarksTableSQL = `CREATE TABLE IF NOT EXISTS checkmarks (
  	MessageId STRING NOT NULL,
  	ConversationId STRING NOT NULL,
  	UserId STRING NOT NULL,
  	Primary Key(MessageId, ConversationId, UserId),
  	CONSTRAINT fk_checkmarks
    	FOREIGN KEY (MessageId, ConversationId) REFERENCES message(MessageId, ConversationId)
      	ON DELETE CASCADE,
    	FOREIGN KEY (UserId) REFERENCES user(UserId)
      	ON DELETE CASCADE
);`

/*
* COMMENT TABLE
*/
var commentTableSQL = `CREATE TABLE IF NOT EXISTS comment (
  	CommentId STRING NOT NULL,
  	UserId STRING NOT NULL,
  	Comment STRING,
  	MessageId STRING NOT NULL,
  	ConversationId STRING NOT NULL,
  	PRIMARY KEY(CommentId, MessageId, ConversationId),
  	CONSTRAINT fk_comment
    	FOREIGN KEY (MessageId, ConversationId) REFERENCES message(MessageId, ConversationId)
      	ON DELETE CASCADE,
    	FOREIGN KEY (UserId) REFERENCES user(UserId)
      	ON DELETE CASCADE
);`
