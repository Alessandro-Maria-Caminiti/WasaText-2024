package database

import (
	"github.com/google/uuid" // Importiamo il pacchetto per gli UUID
	"io"
	"os"
	"fmt"
	"wasatext/service/api/images"
)

// Query per aggiungere un nuovo utente
var Adduser = "INSERT INTO user (UserId, Username, profile_image) VALUES (?, ?, ?);"

func (db *appdbimpl) CreateUser(u User) (User, error) {
    var user User
    user.Username = u.Username
    user.UserId = uuid.New().String()

    userDir := fmt.Sprintf("./storage/profiles/%s", user.UserId)
    conversationsDir := fmt.Sprintf("%s/conversations", userDir)

    fmt.Println("🔧 Creazione cartelle utente:", userDir)
    
    if err := os.MkdirAll(conversationsDir, os.ModePerm); err != nil {
        fmt.Println("❌ ERRORE: impossibile creare la cartella:", conversationsDir, err)
        return user, fmt.Errorf("error creating user directories: %w", err)
    }
    
    fmt.Println("✅ Cartelle create con successo!")

    profilePhotoPath := images.SetDefaultPhoto(user.UserId)
    fmt.Println("🖼 Salvataggio immagine profilo in:", profilePhotoPath)

    source, err := os.Open("./storage/default_profile_photo.jpg")
    if err != nil {
        fmt.Println("❌ ERRORE: impossibile aprire la foto predefinita:", err)
        return user, fmt.Errorf("error opening default profile photo: %w", err)
    }
    defer source.Close()

    destination, err := os.Create(profilePhotoPath)
    if err != nil {
        fmt.Println("❌ ERRORE: impossibile creare la foto profilo:", err)
        return user, fmt.Errorf("error creating profile photo file: %w", err)
    }
    defer destination.Close()

    _, err = io.Copy(destination, source)
    if err != nil {
        fmt.Println("❌ ERRORE: impossibile copiare la foto profilo:", err)
        return user, fmt.Errorf("error copying profile photo: %w", err)
    }

    fmt.Println("✅ Foto profilo copiata con successo!")
    fotina, err := images.ImageToBase64(profilePhotoPath)
    if err != nil{
        fmt.Println("ERRORE IMPOSSIBILE CONVERTIRE FOTO")
        return user, fmt.Errorf("error converting profile photo: %w", err)
    }
    user.Photo = fotina
    fmt.Println("%v", user.Photo)
    _, err = db.c.Exec(Adduser, user.UserId, user.Username, user.Photo)
    if err != nil {
        fmt.Println("❌ ERRORE: impossibile inserire utente nel database:", err)
        return user, fmt.Errorf("error inserting user into database: %w", err)
    }

    fmt.Println("✅ Utente inserito nel database con successo!")
    
    return user, nil
}
