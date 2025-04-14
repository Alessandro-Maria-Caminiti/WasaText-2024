package database

import (
	"log"
	"errors"
)

// UpdateUserProfileImage aggiorna l'immagine profilo di un utente nel database
// Accetta l'ID utente e il path dell'immagine profilo
func (db *appdbimpl) UpdateUserProfileImage(userId string, imagePath string) error {
	// Prepara la query SQL per aggiornare l'immagine profilo
	query := `UPDATE user SET profile_image = ? WHERE UserId = ?`
	
	// Esegui la query
	result, err := db.c.Exec(query, imagePath, userId)
	if err != nil {
		return err
	}
	log.Printf(userId)
	// Verifica se è stata aggiornata una riga
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	// Se nessuna riga è stata aggiornata, significa che l'utente non esiste
	if rowsAffected == 0 {
		return errors.New("utente non trovato")
	}
	
	return nil
}