package database

var FindUserId = `SELECT UserId, Username, profile_image FROM user WHERE UserId = ?;`

func (db *appdbimpl) GetUserById(UserId string) (User, error) {
	var user User
	err := db.c.QueryRow(FindUserId, UserId).Scan(&user.UserId, &user.Username, &user.Photo)
	return user, err
}
