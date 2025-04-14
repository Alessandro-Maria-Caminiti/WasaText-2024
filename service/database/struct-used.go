package database

// -- Structs used in the database for the user containning the UserID and the username -- //
type User struct {
	UserId   string    `json:"UserId"`
	Username string `json:"username"`
	Photo string `json:photo`
}

// -- Structs for the group -- //
type Group struct {
	ConversationId string `json:ConversationId` 
	GroupId    string    `json:"GroupId "`
	GroupName string `json:"groupName"`
	GrouPhoto string `json:grouphoto`
}
