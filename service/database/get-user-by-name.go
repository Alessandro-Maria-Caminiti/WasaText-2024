package database

// Query for search the user by username
var FindbyUsername = `SELECT UserId, Username FROM user WHERE Username = ?;`

func (db *appdbimpl) GetUserByName(username string) (User, error) {
	var user User
	err := db.c.QueryRow(FindbyUsername, username).Scan(&user.UserId, &user.Username)
	return user, err
}
