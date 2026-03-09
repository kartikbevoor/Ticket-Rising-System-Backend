package admin

import (
	database "Ticket_Rising_Backend/Database"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func CreateReply(w http.ResponseWriter, r *http.Request) {
	ticketIdStr := r.URL.Query().Get("id")

	ticketId, err := strconv.Atoi(ticketIdStr)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}
	var reply database.Reply

	err = json.NewDecoder(r.Body).Decode(&reply)

	if err != nil {
		log.Fatal(err)
	}

	database.InsertIntoReplies(&reply, ticketId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reply)
}
