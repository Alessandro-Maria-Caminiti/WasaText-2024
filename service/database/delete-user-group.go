package database

import (
	"database/sql"
)

// Query used do delete in the user_group table
var removemember = `DELETE FROM user_group WHERE GroupId  = ?`

func (db *appdbimpl) DeleteMember(GroupId  string, tx *sql.Tx) error {
	if tx == nil {
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
	}
	// Delete the group in the user table
	_, err := tx.Exec(removemember, GroupId )
	if err != nil {
		return err
	}

	return nil
}
