package database

import (
	"errors"
)

// DeleteUserFromGroup rimuove un utente da un gruppo
func (db *appdbimpl) DeleteUserFromGroup(UserId string, GroupId  string) error {
	// Verifica se il gruppo esiste
	exists, err := db.GroupExists(GroupId )
	if err != nil {
		return errors.New("errore durante la verifica del gruppo")
	}
	if !exists {
		return errors.New("il gruppo non esiste")
	}

	// Esegui la query per rimuovere l'utente dal gruppo
	result, err := db.c.Exec("DELETE FROM user_group WHERE UserId = ? AND GroupId  = ?", UserId, GroupId )
	if err != nil {
		return errors.New("errore durante la rimozione dell'utente dal gruppo")
	}

	// Controlla se è stata effettivamente eliminata una riga
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("l'utente non fa parte del gruppo")
	}

	return nil
}