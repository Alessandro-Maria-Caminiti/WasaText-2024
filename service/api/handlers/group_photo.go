package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"wasatext/service/database"
	"github.com/gorilla/mux"
    "wasatext/service/api/images"
	"io"
    
    "os"
)


// ChangeGroupPhotoHandler handles changing the group photo
func ChangeGroupPhotoHandler(db database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    // Extract sourceChatId from the URL path
    vars := mux.Vars(r)
    sourceChatId := vars["sourceChatId"]
    
    // Convert sourceChatId to integer
    
    // Parse the multipart form (max 10MB)
    err := r.ParseMultipartForm(10 << 20) // 10MB limit
    if err != nil {
        http.Error(w, "Unable to parse form", http.StatusBadRequest)
        return
    }
    
    // Retrieve the file from the form
    file, fileHeader, err := r.FormFile("photo")
    if err != nil {
        http.Error(w, "Unable to retrieve photo", http.StatusBadRequest)
        return
    }
    defer file.Close()
    // Create user-specific directory if not exists
    uploadDir := fmt.Sprintf("./storage/group/%s/", sourceChatId)
    if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
        http.Error(w, "Failed to create user directory", http.StatusInternalServerError)
        return
    }
    // Validate the file type (JPEG, PNG)
    ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
    if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
        http.Error(w, "Invalid file format. Only JPEG and PNG are allowed.", http.StatusBadRequest)
        return
    }
    		// Save the file
		filePath := filepath.Join(uploadDir, "grouphoto"+ext)
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
	PhotoURL := fmt.Sprintf("/uploads/groups/%s/grouphoto%s", sourceChatId, ext)
   
    
    // Update the group photo URL in the database
    err = db.SetGroupPhoto(sourceChatId, PhotoURL)
    if err != nil {
        log.Printf("Error updating group photo: %v", err)
        http.Error(w, "Failed to update group photo", http.StatusInternalServerError)
        return
    }
    
    // Log and respond with success
    log.Printf("Group %s updated photo: %s", sourceChatId, PhotoURL)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status": "Group photo updated successfully",
        "PhotoURL": PhotoURL,
        "Base64": base64Str, 
    })
}
}