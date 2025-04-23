package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"wasatext/service/database"
	"wasatext/service/api/structs"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// SendMessageHandler gestisce l'invio di un messaggio a una conversazione
func SendMessageHandler(database database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	// Estrai l'ID della conversazione dall'URL
	vars := mux.Vars(r)
	conversationId := vars["sourceChatId"]
	// Parse form data
	if err := r.ParseMultipartForm(10 << 20); err != nil { // Limite upload 10MB
		respondWithErrorsndmsg(w, http.StatusBadRequest, "Errore nella richiesta multipart")
		return
	}

	// Estrai i dati dal form
	senderID := vars["id"]
	if senderID == "" {
		respondWithErrorsndmsg(w, http.StatusBadRequest, "Sender ID is required")
		return
	}

	user, err := database.GetUserById(senderID) 
	if err != nil{
		respondWithErrorsndmsg(w, http.StatusInternalServerError, "Errore nella ricerca di chi ha mandato il messaggio")
		return
	}

	content := r.FormValue("content")
	var filePath string // Definiamo la variabile filePath prima dell'uso

	// Gestione dell'upload file
	file, handler, err := r.FormFile("file")
	if err == nil { // Se non c'è errore, significa che è stato caricato un file
		defer file.Close()
		filePath, err = saveUploadedFilesndmsg(file, handler)
		if err != nil {
			respondWithErrorsndmsg(w, http.StatusInternalServerError, "Errore nel salvataggio del file")
			return
		}
	}

	// Verifica che ci sia almeno un contenuto (testo o file)
	if content == "" && filePath == "" {
		respondWithErrorsndmsg(w, http.StatusBadRequest, "Messaggio vuoto")
		return
	}


	log.Printf(user.Username)

	// Crea la struttura del messaggio
	newMessage := structs.Message{
		SenderUserId:   senderID,
		SenderUserName: user.Username,
		ConversationId: conversationId,
		Text:           content,
		Photo:          filePath, // Se il messaggio contiene un file, salviamo il percorso
		SendTime:       time.Now(),
		Status:         "sent", // Imposta lo stato del messaggio
	}
	log.Printf(newMessage.Text)
	// Salva il messaggio nel database
	savedMessage, err := database.CreateMessage(newMessage)
	if err != nil {
		log.Printf("Errore nella creazione del messaggio: %v", err)
		respondWithErrorsndmsg(w, http.StatusInternalServerError, "Errore nell'invio del messaggio")
		return
	}
	log.Printf("%v",savedMessage)

	// Rispondi con il messaggio salvato
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedMessage)
}
}

// SaveUploadedFile salva un file caricato in una directory specificata
func saveUploadedFilesndmsg(file multipart.File, handler *multipart.FileHeader) (string, error) {
	dir := "uploads/"
	os.MkdirAll(dir, os.ModePerm) // Assicura che la cartella esista

	// Genera un nome univoco per il file
	fileName := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(handler.Filename))
	filePath := filepath.Join(dir, fileName)

	outFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

// utils.respondWithError invia una risposta JSON di errore
func respondWithErrorsndmsg(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
