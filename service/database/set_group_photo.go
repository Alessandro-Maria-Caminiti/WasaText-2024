package database

// Query for updating the group photo
var UpdateGroupPhoto = "UPDATE group_t SET PhotoURL = ? WHERE GroupId  = ?;"

// SetGroupPhoto updates the photo URL of a group in the database
func (db *appdbimpl) SetGroupPhoto(GroupId  string, photoUrl string) error {
    // Update the photo URL of the group
    
    _, err := db.c.Exec(UpdateGroupPhoto, photoUrl, GroupId )
    if err != nil {
        return err
    }
    return nil
}