package database

import (
	"database/sql"
)

// Query for delete the user from a conversation
var deleteuserconv = "DELETE FROM conversation_user WHERE UserId = ? AND ConversationId = ?;"

// Function used to remove a user from a conversation
func (db *appdbimpl) DeleteUserConv(UserId string, ConversationId string) error {
	tx, err := db.c.BeginTx(db.ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			err = tx.Rollback()
		}
		err = tx.Commit()
	}()

	// Exec query
	_, err = tx.Exec(deleteuserconv, UserId, ConversationId)
	if err != nil {
		return err
	}

	return nil
}
