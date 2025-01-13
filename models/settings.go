package models

// Settings represents the user settings
type Settings struct {
	Theme         string `json:"theme"`
	Notifications bool   `json:"notifications"`
}
