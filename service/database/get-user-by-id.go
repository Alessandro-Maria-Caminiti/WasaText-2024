package database

var FindUserId = `SELECT UserId, Username FROM user WHERE UserId = ?;`

func (db *appdbimpl) GetUserById(UserId string) (User, error) {
	var user User
	err := db.c.QueryRow(FindUserId, UserId).Scan(&user.UserId, &user.Username)
	return user, err
}
