package database

import "log"

type Ticket struct {
	Id            int    `json:"id"`
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
		FOREIGN KEY (department_id) REFERENCES department(id) ON DELETE CASCADE ON UPDATE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
	)ENGINE=InnoDB;`

	_, err := Db.Exec(ticketTable)
	if err != nil {
		log.Fatal("Error creating ticket table:", err)
	}
}

// This inserts ticket into tickets table
func InsertIntoTickets(ticket *Ticket, userId int) {
	result, err := Db.Exec(
		"INSERT INTO tickets(department_id,description,user_id) VALUES(?,?,?)",
		ticket.Department_Id, ticket.Description, userId,
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

// Fetches tickets from database
func FetchTickets(userId int) []Ticket {

	rows, err := Db.Query(
		"SELECT department_id, description FROM tickets WHERE user_id = ?",
		userId,
	)
	if err != nil {
		log.Println("Unable to fetch tickets:", err)
		return nil
	}
	defer rows.Close()

	var tickets []Ticket

	for rows.Next() {
		var ticket Ticket

		err := rows.Scan(
			&ticket.Department_Id,
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
