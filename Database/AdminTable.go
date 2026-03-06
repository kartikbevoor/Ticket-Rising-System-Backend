package database

import (
	"log"
)

type Admin struct {
	Name          string `json:"name"`
	Id            int    `json:"id"`
	Department_Id int    `json:"department_id"`
	IsSuperAdmin  bool   `json:"is_super_admin"`
	UserName      string `json:"username"`
	Password      string `json:"password"`
}

func CreateAdminTable() {
	adminTable := `
	CREATE TABLE IF NOT EXISTS admin (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100),
		department_id INT,
		is_super_admin BOOLEAN,
		username VARCHAR(100) UNIQUE,
		password VARCHAR(100),
		FOREIGN KEY(department_id) REFERENCES department(id) ON DELETE CASCADE ON UPDATE CASCADE
	)ENGINE=InnoDB;`
	_, err := Db.Exec(adminTable)
	if err != nil {
		log.Fatal("Error creating admin table: ", err)
	}
}

func InsertIntoAdmin(admin Admin) {
	result, err := Db.Exec(
		"INSERT INTO admin(name, department_id, username, password) VALUES(?,?,?)",
		admin.Name, admin.Department_Id, admin.UserName, admin.Password,
	)

	if err != nil {
		log.Fatal("Unable to insert into admin", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Unable to get inserted id:", err)
	}

	admin.Id = int(id)
}
