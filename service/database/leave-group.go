package database

// Function used to remove a member from a group
func (db *appdbimpl) LeaveGroup(UserId string, GroupId  string) error {
	_, err := db.c.Exec("DELETE FROM user_group WHERE GroupId  = ? AND UserId = ?", GroupId , UserId)
	if err != nil {
		return err
	}
	return nil
}
