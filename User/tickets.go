package user

import (
	database "Ticket_Rising_Backend/Database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// This is to raise a ticket
func CreateTicket(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	var ticket database.Ticket

	err = json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		log.Fatal(err)
	}

	// Fields must be validated before inserting
	database.InsertIntoTickets(&ticket, userId) // call function to insert into tickets

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ticket)
}

// This is to view tickets
func ViewTickets(w http.ResponseWriter, r *http.Request) {

	userIdStr := r.URL.Query().Get("id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	tickets := database.FetchTickets(userId) // Fetches from db

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

// This is to view replies
func Viewreplies(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Fatal(err)
	}

	replies := database.FetchReplies(userId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(replies)
}
