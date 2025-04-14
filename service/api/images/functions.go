package images

import (
	"encoding/base64"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

// Function return the path of the user photo profile
func SetDefaultPhoto(userId string) string {
	return fmt.Sprintf("./storage/profiles/%s/default_profile_photo.jpg", userId)
}

// Function return the path of the group photo
func SetDefaultPhotoGroup(groupId string) string {
	return fmt.Sprintf("./storage/groups/%s/default_profile_photo.jpg", groupId)
}
// Convert an image to Base64
func ImageToBase64(filename string) (string, error) {
	imageFile, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() {
		if cerr := imageFile.Close(); cerr != nil {
			fmt.Printf("Errore chiusura file: %v\n", cerr)
		}
	}()

	imageData, err := io.ReadAll(imageFile)
	if err != nil {
		return "", err
	}

	base64Str := base64.StdEncoding.EncodeToString(imageData)
	return base64Str, nil
}

// Save and crop an image
func SaveAndCrop(filename string, w uint, h uint) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("Errore chiusura file: %v\n", cerr)
		}
	}()

	// Decode image
	img, format, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Errore decodifica immagine: %v\n", err)
		return err
	}
	fmt.Printf("Formato immagine: %s\n", format)

	// Resize image
	resizedImg := resize.Resize(w, h, img, resize.NearestNeighbor)

	// Save image
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := out.Close(); cerr != nil {
			fmt.Printf("Errore chiusura file output: %v\n", cerr)
		}
	}()

	// Save in correct format
	if format == "png" {
		err = png.Encode(out, resizedImg)
	} else {
		err = jpeg.Encode(out, resizedImg, nil)
	}
	if err != nil {
		fmt.Printf("Errore salvataggio immagine: %v\n", err)
	}
	return err
}