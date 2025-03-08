// package config

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	_ "github.com/go-sql-driver/mysql"
// )

// var DB *sql.DB

// func ConnectDB() *sql.DB {
// 	var err error

// 	DB, err = sql.Open("mysql", "root:%40dk@tcp(127.0.0.1:3306)/ecommerce")

// 	if err != nil {
// 		log.Fatal("Error connecting to database:", err)
// 	}

// 	// Check database connection
// 	err = DB.Ping()
// 	if err != nil {
// 		log.Fatal("Error pinging database:", err)
// 	}

// 	fmt.Println("Database Connected Successfully")
// 	return DB
// }
