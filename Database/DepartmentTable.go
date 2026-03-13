package database

import (
	"log"
)

type Department struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func CreateDepartmentTable() {
	createTable := `
	CREATE TABLE IF NOT EXISTS department (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100) NOT NULL
	);`

	_, err := Db.Exec(createTable)
	if err != nil {
		log.Fatal("Error creating department table:", err)
	}

	var count int
	err = Db.QueryRow("SELECT COUNT(*) FROM department").Scan(&count)
	if err != nil {
		log.Fatal("Error counting department rows:", err)
	}

	if count == 0 {
		initialDepartments := []string{"Accountant", "Accounts", "CTP", "Design", "Driver", "Press", "ProductionManager", "ProductionStaff"}

		for _, name := range initialDepartments {
			_, err := Db.Exec("INSERT INTO department (name) VALUES (?)", name)
			if err != nil {
				log.Println("Error inserting department:", err)
			}
		}
		log.Println("Department table seeded successfully")
	} else {
		log.Println("Department table already has data; seeding skipped")
	}
}
