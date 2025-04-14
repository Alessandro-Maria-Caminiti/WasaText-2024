package database
import(
    "fmt"
)
// Query used to add the user and the group in the user_group table
var AddtoGroup = "INSERT INTO user_group (UserID, GroupId ) VALUES (?, ?)"

func (db *appdbimpl) AddUserGroup(UserId string, GroupId string) error {
    // Controlla che UserId esista
    var userExists bool
    err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM user WHERE UserId = ?)", UserId).Scan(&userExists)
    if err != nil {
        return fmt.Errorf("error checking if user exists: %w", err)
    }
    if !userExists {
        return fmt.Errorf("user does not exist: %s", UserId)
    }

    // Controlla che GroupId esista
    var groupExists bool
    err = db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM group_t WHERE GroupId = ?)", GroupId).Scan(&groupExists)
    if err != nil {
        return fmt.Errorf("error checking if group exists: %w", err)
    }
    if !groupExists {
        return fmt.Errorf("group does not exist: %s", GroupId)
    }

    // Inserisci nella tabella user_group
    _, err = db.c.Exec(AddtoGroup, UserId, GroupId)
    if err != nil {
        return fmt.Errorf("error inserting into user_group: %w", err)
    }

    return nil
}
