package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"wasatext/service/api/images"
	"wasatext/service/api/models"
	"wasatext/service/database"
)

// LoginHandler handles user login and creation.
func LoginHandler(db database.AppDatabase) http.HandlerFunc {
	fmt.Printf("LoginHandler factory: db is nil? %v\n", db == nil)
	return func(w http.ResponseWriter, r *http.Request) {
	var userRequest models.UserRequest
	
	if db == nil {
		log.Fatal("Database connection is nil")
	}

	// Decodifica il body JSON nella struttura userRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Verifica che sia stato fornito un Nome
	if userRequest.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	// Controlla se l'utente esiste nel database cercando per Nome
	var userId string
	var exists bool
	log.Printf("Cerco di prendere lo User?")
	user, err := db.GetUserByName(userRequest.Name)
	if err != nil {
		log.Printf("sono io il problema?")
		if err.Error() == "sql: no rows in result set" {
			log.Printf("Hallo")
			exists = false
		} else {
			log.Printf("Database error in GetUserByName: %v", err)
			http.Error(w, "Error checking user", http.StatusInternalServerError)
			return
		}
	} else {
		log.Printf("lo riesco a fare?")
    	exists = true
    	userId = user.UserId // ⚠️ Salviamo l'ID per il successivo GetUserById!
	}
	// Se l'utente non esiste, crealo
	if !exists {

		newUser := database.User{
			Username: userRequest.Name, // Usare il Nome come Username
		}
		log.Printf("STO CREANDO L'USER")
		createdUser, err := db.CreateUser(newUser)
		if err != nil {
			log.Printf("Error creating user: %v", err)
			http.Error(w, "Error during registration", http.StatusInternalServerError)
			return
		}

		userId = createdUser.UserId
	}
	userphoto, err := images.ImageToBase64(images.SetDefaultPhoto(userId))
	if err != nil{
		http.Error(w,"Error setting Default Photo", http.StatusInternalServerError)
	}
	// Ottieni informazioni complete dell'utente
	user, err = db.GetUserById(userId)
	log.Printf(userId)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		http.Error(w, "Error retrieving user data", http.StatusInternalServerError)
		return
	}

	// Risposta con le informazioni dell'utente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "Login successful",
		"user":   user,
		"userId": userId, 
		"photo": userphoto,
	})
}
}
