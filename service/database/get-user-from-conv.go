package database

import (
	_"github.com/mattn/go-sqlite3"
	"log"
)

var getuserz = `SELECT user.UserId, Username FROM conversation_user JOIN user ON conversation_user.UserId = user.UserId WHERE ConversationId = ?`

func (db *appdbimpl) SearchUsersfromconv(convid string) ([]User, error) {
	var users []User
	log.Printf(convid)
	rows, err := db.c.Query(getuserz, convid)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	log.Printf("%b", rows)
	for rows.Next() {
		if rows.Err() != nil {
			return nil, err
		}

		var u User
		if err := rows.Scan(&u.UserId, &u.Username); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	defer func() { err = rows.Close() }()

	return users, err
}
