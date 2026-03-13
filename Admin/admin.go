package admin

import (
	database "Ticket_Rising_Backend/Database"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func AdminStarting() {

}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {

	_, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var admin database.Admin

	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.InsertIntoAdmin(admin)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(admin)
}

func VerifyAdminCredentials(w http.ResponseWriter, r *http.Request) {

	_, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	isValidAdmin := database.MatchAdminCredentials(username, password)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(isValidAdmin)
}
