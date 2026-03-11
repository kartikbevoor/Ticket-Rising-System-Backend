package database

import (
	"log"
)

type Reply struct {
	Id        int    `json:"id"`
	Ticket_Id int    `json:"ticket_id"`
	Comment   string `json:"comment"`
}

func CreateRepliesTable() {
	repliesTable := `
	CREATE TABLE IF NOT EXISTS replies(
		id INT PRIMARY KEY AUTO_INCREMENT,
		comment VARCHAR(200) NOT NULL,
		ticket_id INT,
		FOREIGN KEY(ticket_id) REFERENCES tickets(id) ON DELETE CASCADE ON UPDATE CASCADE
	)ENGINE=InnoDB;`

	_, err := Db.Exec(repliesTable)
	if err != nil {
		log.Fatal("Error creating user table: ", err)
	}
}

func InsertIntoReplies(reply *Reply, ticketId int) {
	result, err := Db.Exec(
		"INSERT INTO replies(ticket_id, comment) VALUES(?,?)",
		ticketId, reply.Comment,
	)

	if err != nil {
		log.Fatal("Unable to insert into replies", err)
	}

	UpdateTicketStatusToResolved(ticketId)

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Unable to get inserted id:", err)
	}

	reply.Id = int(id)
}

type responseReply struct {
	Comment string `json:"comment"`
}

func FetchReplies(userId int) []responseReply {

	rows, err := Db.Query(
		`SELECT comment
		 FROM replies
		 WHERE ticket_id IN (
			 SELECT id FROM tickets WHERE user_id = ?
		 )`,
		userId,
	)

	if err != nil {
		log.Println("Unable to fetch replies:", err)
		return nil
	}
	defer rows.Close()

	var replies []responseReply

	for rows.Next() {
		var reply responseReply

		err := rows.Scan(
			&reply.Comment,
		)

		if err != nil {
			log.Println("Scan error:", err)
			continue
		}

		replies = append(replies, reply)
	}

	return replies
}
