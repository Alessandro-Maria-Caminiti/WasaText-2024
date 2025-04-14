package database

import (
    "github.com/google/uuid"
	"io"
	"os"
    "strings"
    "fmt"
    "log"
	"wasatext/service/api/images"
    "wasatext/service/api/structs"
)

// Query for create a group in the group table
var Addgroup = "INSERT INTO group_t (ConversationId,GroupId , GroupName, GrouPhoto) VALUES (?, ?, ?, ?)"

func (db *appdbimpl) CreateGroup(GroupName string, UserId string) (Group, error) {
    var g Group
    var conversation structs.Conversation
    g.GroupName = GroupName
    
	newConv, err := db.CreateConversation(conversation)
	if err != nil{
		return g, err
	}
    err = db.AddUserConv(newConv.ConversationId, UserId)
    if err != nil {
        return g, err
    }

    g.ConversationId = newConv.ConversationId

    // Verifica che UserId non sia vuoto
    if UserId == "" {
        return g, fmt.Errorf("invalid UserId: UserId is empty")
    }

    // Genera un nuovo UUID per il gruppo
    g.GroupId = uuid.New().String()

    // Creazione della cartella del gruppo
    path := "./storage/groups/" + g.GroupId
    if err := os.MkdirAll(path, os.ModePerm); err != nil {
        return g, err
    }

    // Creazione della foto profilo del gruppo
    source, err := os.Open("./storage/default_profile_photo.jpg")
    if err != nil {
        return g, err
    }
    destination, err := os.Create(images.SetDefaultPhotoGroup(g.GroupId)) // Create the path where the photo will be saved
	if err != nil {
		return g, err
	}
	defer destination.Close() // Close the user folder

    _, err = io.Copy(destination, source)
    if err != nil {
        return g, err
    }
    var buf strings.Builder
    _, err = io.Copy(&buf, source)
    if err != nil {
    return g, err
    }
    s := buf.String()
    // Inserimento nel db
    _, err = db.c.Exec(Addgroup,g.ConversationId, g.GroupId, g.GroupName, s)
    if err != nil {
        log.Printf("??????")
        return g, err
    }
    fmt.Printf("Groupname %s\n", g.GroupName)
    // Log per verificare UserId
    fmt.Printf("DEBUG: UserId passato a AddUserGroup: '%s'\n", UserId)

    // Verifica se UserId esiste prima di aggiungerlo
    err = db.AddUserGroup(UserId, g.GroupId)
    if err != nil {
        return g, fmt.Errorf("failed to add user to group: %w", err)
    }
    return g, nil
}