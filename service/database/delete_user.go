package database
import(
    "log"
)
// Query for deleting a user by ID
var DeleteUserById = "DELETE FROM user WHERE UserId = ?;"

// DeleteUser removes a user from the database by their ID
func (db *appdbimpl) DeleteUser(UserId string) error {
    // Start a transaction since we might need to delete related data
    tx, err := db.c.Begin()
    if err != nil {
        return err
    }
    
    // Defer a rollback in case anything fails
    defer func() {
        if err != nil {
            tx.Rollback()
        }
    }()
    
    // First, delete user's group memberships
    _, err = tx.Exec("DELETE FROM user_group WHERE UserId = ?", UserId)
    if err != nil {
        log.Printf("UserGroup error:",err.Error())
        return err
    }
    
    // Delete user from conversation_user table
    _, err = tx.Exec("DELETE FROM conversation_user WHERE UserId = ?", UserId)
    if err != nil {
        log.Printf("Conversation error:",err.Error())
        return err
    }
    
    // Delete user's comments
    _, err = tx.Exec("DELETE FROM comment WHERE UserId = ?", UserId)
    if err != nil {
        log.Printf("Comments error:",err.Error())
        return err
    }
    
    // Delete checkmarks associated with the user
    _, err = tx.Exec("DELETE FROM checkmarks WHERE UserId = ?", UserId)
    if err != nil {
        log.Printf("Checkmarks error:",err.Error())
        return err
    }
    
    // Finally, delete the user
    _, err = tx.Exec(DeleteUserById, UserId)
    if err != nil {
        log.Printf("Eliminazione error:",err.Error())
        return err
    }
    
    // Commit the transaction
    err = tx.Commit()
    if err != nil {
        return err
    }
    
    return nil
}