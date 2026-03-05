package database

import "log"

type Department struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func CreateDepartmentTable() {
	departmentTable := `
	CREATE TABLE IF NOT EXISTS department(
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100)
	);`
	_, err := Db.Exec(departmentTable)
	if err != nil {
		log.Fatal("Error creating category table", err)
	}
}
