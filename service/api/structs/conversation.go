package structs

/*
* Struct used to rapresent a conversation
* if the groupId is 0 the conversation is with a user
* if the groupId is different from 0 the conversation is with a groupa
* lastMessageId is the id of the last message of the conversation */

type Conversation struct {
	ConversationId string `json:"conversationId"`
	GroupId        string `json:"GroupId"`
	LastMessageId  string `json:"lastMessageId"`
}
