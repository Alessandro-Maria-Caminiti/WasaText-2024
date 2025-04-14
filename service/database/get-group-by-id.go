package database

var FindGroupId  = `SELECT GroupId , GroupName FROM group_t WHERE GroupId  = ?;`

func (db *appdbimpl) GetGroupById(GroupId  string) (Group, error) {
	var group Group
	err := db.c.QueryRow(FindGroupId , GroupId ).Scan(&group.GroupId , &group.GroupName)
	return group, err
}
