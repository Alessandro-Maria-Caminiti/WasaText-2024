package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"mime/multipart"
	"github.com/google/uuid"
	"wasatext/service/api/images"
	"wasatext/service/database"
)

// PhotoUploadResponse represents the response sent back to the client
type PhotoUploadResponse struct {
	Status   string `json:"status"`
	PhotoURL string `json:"photoUrl"`
	Photo   string `json:"photo"` // Base64 opzionale
}

// SetMyPhotoHandler handles updating a user's profile photo
func SetMyPhotoHandler(database database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract user ID from URL
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 4 || parts[2] == "" {
			http.Error(w, "Missing user ID", http.StatusBadRequest)
			return
		}
		userID := parts[2]

		// Parse form data (limit upload size to 10MB)
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		// Retrieve the file
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Invalid file upload", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Create user-specific directory if not exists
		uploadDir := fmt.Sprintf("./storage/users/%s/", userID)
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			http.Error(w, "Failed to create user directory", http.StatusInternalServerError)
			return
		}

		// Validate and determine file extension
		ext := strings.ToLower(filepath.Ext(handler.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			http.Error(w, "Invalid file format. Only JPEG and PNG are allowed.", http.StatusBadRequest)
			return
		}

		// Save the file
		filePath := filepath.Join(uploadDir, "profile"+ext)
		outFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Unable to save the photo", http.StatusInternalServerError)
			return
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, file)
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}
		
		// Convert image to Base64
		base64Str, err := images.ImageToBase64(filePath)
		if err != nil {
			log.Printf("Base64 conversion error: %v", err)
			base64Str = "" // If conversion fails, leave it empty
		}

		// Resize image to 300x300 pixels
		err = images.SaveAndCrop(filePath, 300, 300)
		if err != nil {
			log.Printf("Image resize error: %v", err)
		}

		// Generate the public URL
		photoURL := fmt.Sprintf("/uploads/users/%s/profile%s", userID, ext)

		// Update database with the new image path
		err = database.UpdateUserProfileImage(userID, base64Str)
		if err != nil {
			log.Printf(err.Error())
			os.Remove(filePath) // Delete file if DB update fails
			http.Error(w, "Failed to update user profile in database", http.StatusInternalServerError)
			return
		}
		var response PhotoUploadResponse
		response.Status = "user Photo Updated Successfully"
		response.PhotoURL= photoURL
		response.Photo = base64Str
		// Send JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

// utils.respondWithError invia una risposta JSON di errore
func respondWithErrorusrpht(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// SaveUploadedFile salva un file caricato in una directory specificata
func saveUploadedFileusrpht(file multipart.File, handler *multipart.FileHeader) (string, error) {
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

// isValidImageFile checks if the file has a valid image extension
func isValidImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

