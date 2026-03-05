package main

import (
	database "Ticket_Rising_Backend/Database"
	user "Ticket_Rising_Backend/User"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db() // Database
	mux := http.NewServeMux()
	mux.HandleFunc("/createUserAccount", user.CreateUser)    // to create user account
	mux.HandleFunc("/userLogin", user.VerifyUserCredentials) // user login
	mux.HandleFunc("/raiseTicket", user.CreateTicket)        // raise a ticket
	mux.HandleFunc("/viewTickets", user.ViewTickets)         // view tickets
	mux.HandleFunc("/viewReplies", user.Viewreplies)         // view replies
	http.ListenAndServe(":8080", mux)

	log.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}

func db() {
	database.CheckIfDbExists() // checks if db exists if not creates one
	database.ConnectDb()       // connects to the db

	database.CreateUserTable()       // creats user table if it does not exist
	database.CreateDepartmentTable() // creats department table if it does not exist
	database.CreateAdminTable()      // creats admin table if it does not exist
	database.CreateTicketTable()     // creats ticket table if it does not exist
	database.CreateRepliesTable()    // creats replies table if it does not exist
	fmt.Println("Created and connected to database")
}
