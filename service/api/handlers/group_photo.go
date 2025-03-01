package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

// ChangeGroupPhotoHandler handles changing the group photo
func ChangeGroupPhotoHandler(w http.ResponseWriter, r *http.Request) {
	// Extract sourceChatId from the URL path
	vars := mux.Vars(r)
	sourceChatId := vars["sourceChatId"]

	// Parse the multipart form (max 10MB)
	err := r.ParseMultipartForm(10 << 20) // 10MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Retrieve the file and description from the form
	file, fileHeader, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Unable to retrieve photo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate the file type (JPEG, PNG)
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename)) // Use fileHeader.Filename to get the filename
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		http.Error(w, "Invalid file format. Only JPEG and PNG are allowed.", http.StatusBadRequest)
		return
	}

	// Simulate saving the photo (in a real scenario, save to storage and generate a URL)
	photoUrl := fmt.Sprintf("https://example.com/images/%s-group-photo%s", sourceChatId, ext)

	// Log and respond with success
	log.Printf("Group %s updated photo: %s", sourceChatId, photoUrl)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   "Group photo updated successfully",
		"photoUrl": photoUrl,
	})
}
