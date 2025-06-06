package database

import (
	"errors"
	"fmt"
)
// AddToGroup adds a user to a group in the database.
// It first checks if the group name already exists as a username, and if so, returns an error.

func (db *appdbimpl) AddUserToGroup(groupname string, username string) error {

	// Check if the groupname already exists as a username
	_, err := db.GetUser(groupname)
	if err == nil {
		return fmt.Errorf("group name already exists as a username")
	} else if !errors.Is(err, ErrUserNotFound) {
		return fmt.Errorf("database error while checking username: %w", err)
	}

	// Check if the group exists
	var groupCount int
	err = db.c.QueryRow(`
		SELECT COUNT(*) 
		FROM groups 
		WHERE groupname = ?;
	`, groupname).Scan(&groupCount)
	var userAlreadyMember int
	err = db.c.QueryRow(`
		SELECT COUNT(*) 
		FROM group_members 
		WHERE groupname = ? AND membername = ?;
	`, groupname, username).Scan(&userAlreadyMember)

	if err != nil {
		return fmt.Errorf("failed to check if user is already a member: %w", err)
	}

	if userAlreadyMember > 0 {
		return fmt.Errorf("user is already a member: %w", err)	// Skip adding the user if they are already a member
	}

	// Add the user to the group
	_, err = db.c.Exec(`
		INSERT INTO group_members (groupname, membername)
		VALUES (?, ?);
	`, groupname, username)

	if err != nil {
		return fmt.Errorf("failed to add user %s to group: %w", username, err)
	}
	return nil
}