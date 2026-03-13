package user

import (
	database "Ticket_Rising_Backend/Database"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// This is nothing but user registration
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user database.User

	_, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.InsertIntoUser(&user) // Insert user into users table

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Function to verify user
func VerifyUserCredentials(w http.ResponseWriter, r *http.Request) {

	_, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// fmt.Println("VerifyUserCredentials called")

	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	// fmt.Println("Username:", username)
	// fmt.Println("Password:", password)

	isValidUser := database.MatchUserCredentials(username, password) // Checks database to varify user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(isValidUser)
}
