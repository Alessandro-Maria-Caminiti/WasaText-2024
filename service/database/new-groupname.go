package database

// Query for update the username of the user
var UpdateGroupname = "UPDATE group_t SET GroupName = ? WHERE GroupId  = ?;"

func (db *appdbimpl) SetGroupName(GroupId  string, newName string) error {
	// Update the username of the user
	_, err := db.c.Exec(UpdateGroupname, newName, GroupId )
	if err != nil {
		return err
	}
	return nil
}
