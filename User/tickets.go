package user

import (
	database "Ticket_Rising_Backend/Database"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

// This is to raise a ticket
func CreateTicket(w http.ResponseWriter, r *http.Request) {

	_, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

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

	_, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

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
func ViewReplies(w http.ResponseWriter, r *http.Request) {

	_, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	userIdStr := r.URL.Query().Get("id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	replies := database.FetchReplies(userId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(replies)
}

// This is to admin to view tickets so he can reply
func ViewTicketsToAdmin(w http.ResponseWriter, r *http.Request) {

	_, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// fmt.Println("Sever started to get tickets")
	adminIdStr := r.URL.Query().Get("id")

	adminId, err := strconv.Atoi(adminIdStr)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	// need to write code for to check for super admin if s he should view all tickets
	isSuperAdmin := database.CheckIsSuperAdmin(adminId)
	// fmt.Println("Afer check admin type:", isSuperAdmin)

	if isSuperAdmin {
		tickets := database.SuperAdminTickets()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tickets)
		return
	} else {
		tickets := database.FetchAdminTickets(adminId) // Fetches from db

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tickets)
		return
	}
}
