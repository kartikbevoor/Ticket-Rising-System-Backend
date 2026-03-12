package database

import (
	"fmt"
	"log"
)

func CreateStatusTable() {
	statusTable := `CREATE TABLE IF NOT EXISTS status ( 
		id INT PRIMARY KEY AUTO_INCREMENT, 
		type VARCHAR(100) NOT NULL
	)ENGINE=InnoDB;`

	_, err := Db.Exec(statusTable)
	if err != nil {
		log.Fatal("Error creating status table: ", err)
	}

	var count int
	err = Db.QueryRow("SELECT COUNT(*) FROM status").Scan(&count)
	if err != nil {
		log.Fatal("Error counting status rows:", err)
	}

	if count == 0 {
		statuses := []string{"Raised", "Processing", "Resolved"}

		for _, name := range statuses {
			_, err := Db.Exec("INSERT INTO status(type) VALUES (?)", name)
			if err != nil {
				log.Println("Error inserting status:", err)
			}
		}
		log.Println("Status table created and added values")
	} else {
		log.Println("Status table already has data")
	}

	fmt.Println("Created status table")
}
