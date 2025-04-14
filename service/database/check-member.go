package database

import (
	"errors"
	"log"
	"database/sql"
)

// Query used to check if a user is member of a group
var Memeberfromgroup = `SELECT UserId FROM user_group WHERE UserId = ? AND GroupId  = ?`

// Function used to check id a user is member of a group
func (db *appdbimpl) CheckMember(UserId string, GroupId string) (bool, error) {
    var userId string
    err := db.c.QueryRow(Memeberfromgroup, UserId, GroupId).Scan(&userId)
    if err != nil {
        if err == sql.ErrNoRows {
            return false, nil // Nessun record trovato
        }
        log.Printf(err.Error())
        return false, errors.New("Error checking group membership")
    }
    return true, nil // Se abbiamo un risultato, l'utente è un membro
}