package database

import "log"

type Ticket struct {
	Id            int    `json:"id"` // `json:"id,omitempty"`
	Department_Id int    `json:"department_id"`
	Description   string `json:"description"`
	User_Id       int    `json:"user_id"`
}

// This creats tickets table if not exists
func CreateTicketTable() {
	ticketTable := `
	CREATE TABLE IF NOT EXISTS tickets (
		id INT PRIMARY KEY AUTO_INCREMENT,
		department_id INT NOT NULL,
		description TEXT NOT NULL,
		ticket_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		user_id INT,
		status_id int,
		FOREIGN KEY (department_id) REFERENCES department(id) ON DELETE CASCADE ON UPDATE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
		FOREIGN KEY (status_id)	REFERENCES status(id)
	)ENGINE=InnoDB;`

	_, err := Db.Exec(ticketTable)
	if err != nil {
		log.Fatal("Error creating ticket table:", err)
	}
}

// This inserts ticket into tickets table
func InsertIntoTickets(ticket *Ticket, userId int) {
	StatusId := 1
	result, err := Db.Exec(
		"INSERT INTO tickets(department_id,description,user_id,status_id) VALUES(?,?,?,?)",
		ticket.Department_Id, ticket.Description, userId, StatusId,
	)
	if err != nil {
		log.Fatal("Unable insert into tickets", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Unable to get inserted id:", err)
	}

	ticket.Id = int(id)
}

// omitempty
// Creating a new response type for ticket (we can also use omitempty)
type TicketResponseUser struct {
	Department_Id int    `json:"department_id"`
	Ticket_Status string `json:"ticket_status"`
	Description   string `json:"description"`
}

// Fetches tickets from database for user view
// Fetches tickets from database for user view
func FetchTickets(userId int) []TicketResponseUser {

	rows, err := Db.Query(
		`SELECT t.department_id, s.type, t.description
		 FROM tickets t
		 JOIN status s ON t.status_id = s.id
		 WHERE t.user_id = ?`,
		userId,
	)
	if err != nil {
		log.Println("Unable to fetch tickets:", err)
		return nil
	}
	defer rows.Close()

	var tickets []TicketResponseUser

	for rows.Next() {
		var ticket TicketResponseUser

		err := rows.Scan(
			&ticket.Department_Id,
			&ticket.Ticket_Status,
			&ticket.Description,
		)

		if err != nil {
			log.Println("Scan error:", err)
			continue
		}

		tickets = append(tickets, ticket)
	}

	return tickets
}

type TicketResponseAdmin struct {
	// Id          int    `json:"id"`
	Description string `json:"description"`
}

// This is for admin to view tickets for reply
func FetchAdminTickets(adminId int) []TicketResponseAdmin {
	var tickets []TicketResponseAdmin

	rows, err := Db.Query(
		`SELECT t.description 
		 FROM tickets t 
		 WHERE t.department_id = (
			SELECT department_id FROM admin WHERE id = ?)`,
		adminId,
	)

	if err != nil {
		log.Println("Unable to fetch tickets:", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var ticket TicketResponseAdmin
		err := rows.Scan(&ticket.Description)
		if err != nil {
			log.Println("Scan error:", err)
			continue
		}
		tickets = append(tickets, ticket)
	}

	if err = rows.Err(); err != nil {
		log.Println("Rows iteration error:", err)
	}

	return tickets
}

func SuperAdminTickets() []TicketResponseAdmin {
	var tickets []TicketResponseAdmin

	rows, err := Db.Query(
		`SELECT description 
		 FROM tickets`,
	)

	if err != nil {
		log.Println("Unable to fetch tickets:", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var ticket TicketResponseAdmin
		err := rows.Scan(&ticket.Description)
		if err != nil {
			log.Println("Scan error:", err)
			continue
		}
		tickets = append(tickets, ticket)
	}

	if err = rows.Err(); err != nil {
		log.Println("Rows iteration error:", err)
	}

	return tickets
}

func UpdateTicketStatusToResolved(ticketId int) {
	ticketStatus := 3

	_, err := Db.Exec(
		"UPDATE tickets SET status_id = ? WHERE id = ?",
		ticketStatus,
		ticketId,
	)

	if err != nil {
		log.Println("Updating ticket status failed:", err)
	}
}
