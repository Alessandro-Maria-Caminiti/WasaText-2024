package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	//"log"
	"wasatext/service/database"
	"github.com/gorilla/mux"
)

type Userresponse struct {
	UserId   string `json:"userid"`
	Username string `json:"username"`
	Photo   string `json:"photo"` // Base64 opzionale
}


// UserHandler handles user-specific GET requests.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "User endpoint"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// DeleteUserHandler elimina un utente in base all'ID fornito nell'URL. 
func DeleteUserHandler(database database.AppDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    // Per la route settings/{id}/deleteuser
    parts := strings.Split(r.URL.Path, "/")
    
    // Controlliamo che ci siano abbastanza parti nell'URL e che l'ID non sia vuoto
    // L'URL dovrebbe essere del tipo /settings/{id}/deleteuser
    // quindi l'ID si trova nella posizione parts[2]
	//log.Printf(parts[3])
    if len(parts) < 4 || parts[3] == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Missing user ID"})
        return
    }
	userID := parts[3]

	err := database.DeleteUser(userID)
	success := err == nil

	if success {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully", "userID": userID})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to delete user"})
	}
	}
}

func GetUser(database database.AppDatabase) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var User Userresponse
		vars := mux.Vars(r)
		userId := vars["Id"]
		user, err := database.GetUserById(userId)
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Missing User Id or User doesn't exist"})
			return
		}
		User.Username = user.Username
		User.UserId = userId
		User.Photo = user.Photo
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success",
		"user" : User,
	})
}
}