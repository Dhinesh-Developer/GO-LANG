package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbm *sql.DB

func connectDB() {
	var err error
	dbm, err = sql.Open("mysql", "root:root@dk@tcp(127.0.0.1:3306)/jdbc?parseTime=true")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Check if the database connection is successful
	err = dbm.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("MySQL Connected....")
}

func insertInfo() {
	name := "Kumar"
	course := "Go Lang"
	result, err := dbm.Exec(`INSERT INTO go (name, course) VALUES (?, ?)`, name, course)

	if err != nil {
		log.Fatal("Error inserting record:", err)
	} else {
		value, _ := result.LastInsertId()
		fmt.Println("Inserted record ID:", value)
	}
}

func createTable() {
	query := `CREATE TABLE IF NOT EXISTS go (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name TEXT NOT NULL,
		course TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := dbm.Exec(query)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	fmt.Println("Table 'go' created successfully")
}

func updateInfo() {
	id := 3
	name := "DK"
	course := "GO Lang Developer"

	result, err := dbm.Exec(`UPDATE go SET name=?, course=? WHERE id=?`, name, course, id)
	if err != nil {
		log.Fatal("Error updating record:", err)
	} else {
		value, _ := result.RowsAffected()
		fmt.Println("Number of rows updated:", value)
	}
}

func getAll() {
	type Go struct {
		id        int
		name      string
		course    string
		created_at time.Time
	}
	var g Go

	rows, err := dbm.Query("SELECT id, name, course, created_at FROM go")
	if err != nil {
		log.Fatal("Error fetching records:", err)
	}
	defer rows.Close() 

	for rows.Next() {
		err := rows.Scan(&g.id, &g.name, &g.course, &g.created_at)
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}
		fmt.Println("Go ID =", g.id)
		fmt.Println("Go Name =", g.name)
		fmt.Println("Go Course =", g.course)
		fmt.Println("Go Time =", g.created_at)
		fmt.Println("---------------------")
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Error with rows:", err)
	}
}

func deleteInfo(){
	id := 3
	_,err := dbm.Exec(`Delete from go where id = ?`,id)
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("Record deleted id :=" , id)
	}

}

func main() {
	connectDB()
	// createTable()
	// insertInfo()
	// updateInfo()
	getAll()
	deleteInfo()
}
