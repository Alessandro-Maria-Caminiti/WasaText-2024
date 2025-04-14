package models

// UserRequest represents the structure of a user login request.
type UserRequest struct{
	Name string `json:"Name"`
}

type User struct {
	UserId   string    `json:"userId"`
	Username string `json:"username"`
	Photo    string `json:"photo"`
}