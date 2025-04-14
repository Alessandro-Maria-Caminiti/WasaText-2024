package database

// Query
var CheckConv = `
		SELECT COUNT(convp.ConversationId)
		FROM conversation_user convp, conversation_user conv, conversation c
		WHERE convp.ConversationId = conv.ConversationId AND convp.UserId = ? AND conv.UserId = ? 
      AND c.ConversationId = convp.ConversationId AND c.GroupId  IS NULL
	`
var CheckUserConv = `
  SELECT ConversationId
  FROM conversation_user
  WHERE ConversationId = ? AND UserId = ?
`

var GetSenderUserI = `
	SELECT convp.ConversationId
	FROM conversation_user convp, conversation_user conv, conversation c
	WHERE convp.ConversationId = conv.ConversationId AND convp.UserId = ? AND conv.UserId = ? 
	AND c.ConversationId = convp.ConversationId AND c.GroupId  IS NULL
`

func (db *appdbimpl) CheckIfExistConv(sender string, receiver string) (bool, error) {
	var count int
	err := db.c.QueryRow(CheckConv, sender, receiver).Scan(&count)
	if err != nil {
		return false, err
	}

	return count == 1, nil
}

func (db *appdbimpl) CheckUserConv(UserId string, ConversationId string) (bool, error) {
	var id int
	err := db.c.QueryRow(CheckUserConv, ConversationId, UserId).Scan(&id)
	if err != nil {
		return false, err
	}
	return id != 0, nil
}

func (db *appdbimpl) GetConversation(sender string, receiver string) (string, error) {
	var conv string
	err := db.c.QueryRow(GetSenderUserI, sender, receiver).Scan(&conv)
	if err != nil {
		return conv, err
	}
	return conv, nil
}
