package admin

import (
	database "Ticket_Rising_Backend/Database"
	"encoding/json"
	"net/http"
)

func AdminStarting() {

}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
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
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	isValidAdmin := database.MatchAdminCredentials(username, password)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(isValidAdmin)
}
