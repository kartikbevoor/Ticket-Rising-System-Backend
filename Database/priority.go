package database

import (
	"fmt"
	"log"
)

func CreatePriorityTable() {
	priorityTable := `CREATE TABLE IF NOT EXISTS priority ( 
		id INT PRIMARY KEY AUTO_INCREMENT, 
		type VARCHAR(100) NOT NULL
	) ENGINE=InnoDB;`

	_, err := Db.Exec(priorityTable)
	if err != nil {
		log.Fatal("Error creating priority table: ", err)
	}

	var count int
	err = Db.QueryRow("SELECT COUNT(*) FROM priority").Scan(&count)
	if err != nil {
		log.Fatal("Error counting priority rows:", err)
	}

	if count == 0 {
		priorities := []string{"Low", "Medium", "High"}

		for _, name := range priorities {
			_, err := Db.Exec("INSERT INTO priority(type) VALUES (?)", name)
			if err != nil {
				log.Println("Error inserting into priority:", err)
			}
		}
		log.Println("Priority table created and added values")
	} else {
		log.Println("Priority table already has data")
	}
	fmt.Println("created priority table")
}
