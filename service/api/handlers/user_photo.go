package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const photoDir = "./uploads" // Directory where photos will be stored

// SetUserPhoto handles the POST /settings/userphoto route
func SetUserPhoto(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data with a limit of 10MB
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to process the form data", http.StatusInternalServerError)
		return
	}

	// Get the photo file from the form
	file, _, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "No file uploaded", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get the description (optional)
	description := r.FormValue("description")

	// Check the file extension to ensure it's a valid image format (JPEG, PNG)
	ext := strings.ToLower(filepath.Ext(description))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		http.Error(w, "Invalid file format. Only JPEG and PNG are allowed.", http.StatusBadRequest)
		return
	}

	// Generate a unique filename for the photo
	// This can be extended to use more advanced unique identifiers like UUIDs
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(photoDir, fileName)

	// Create the uploads directory if it doesn't exist
	err = os.MkdirAll(photoDir, os.ModePerm)
	if err != nil {
		http.Error(w, "Unable to create upload directory", http.StatusInternalServerError)
		return
	}

	// Save the photo to disk
	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to save the photo", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Copy the uploaded file to the server
	_, err = ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// Respond with the success message and the URL of the uploaded photo
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"status":   "User photo updated successfully",
		"photoUrl": "https://example.com/images/" + fileName, // Replace with actual URL where the photo can be accessed
	}
	if description != "" {
		response["description"] = description
	}
	json.NewEncoder(w).Encode(response)
}
