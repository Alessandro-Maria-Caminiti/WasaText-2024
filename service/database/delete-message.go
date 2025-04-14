package database

import (
	"database/sql"
)

// Query used to remeove a message from the db
var deletemessage = `DELETE FROM message WHERE MessageId = ? AND ConversationId = ?`

func (db *appdbimpl) DeleteMessage(messageId string, ConversationId string) error {
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
	_, err = tx.Exec(deletemessage, messageId, ConversationId)
	if err != nil {
		return err
	}

	return nil
}
