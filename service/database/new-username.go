package database

// Query for update the username of the user
var Updateusername = "UPDATE user SET Username = ? WHERE UserId = ?;"

func (db *appdbimpl) SetMyUsername(UserId string, newUsername string) error {
	// Update the username of the user
	_, err := db.c.Exec(Updateusername, newUsername, UserId)
	if err != nil {
		return err
	}
	return nil
}
