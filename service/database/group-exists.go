package database

import (
	"errors"
)

// GroupExists verifica se un gruppo esiste nel database
func (db *appdbimpl) GroupExists(GroupId  string) (bool, error) {
	var exists bool

	// Query per verificare l'esistenza del gruppo
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM group_t WHERE GroupId  = ?)", GroupId ).Scan(&exists)
	if err != nil {
		return false, errors.New("errore durante la verifica del gruppo")
	}

	return exists, nil
}