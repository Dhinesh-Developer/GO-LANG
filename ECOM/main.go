package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Global database variable
var DB *sql.DB

// Database Connection Function
func ConnectDB() {
	var err error
	dsn := "root:%40dk@tcp(127.0.0.1:3306)/ecommerce" // Corrected DSN with URL encoding for '@'
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// // Check connection
	// if err = DB.Ping(); err != nil {
	// 	log.Fatalf("Database ping failed: %v", err)
	// }

	fmt.Println("✅ Database Connected Successfully")
}

// Product Model
type Product struct {
	ID          int
	Name        string
	Description string
	Price       string
}

// Homepage Handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	// Fetch products from the database
	rows, err := DB.Query("SELECT id, name, description, price FROM products")
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price); err != nil {
			http.Error(w, "Error scanning products", http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	// Parse and execute template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, products)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// Register Routes
func RegisterRoutes() {
	http.HandleFunc("/", HomePage)
}

func main() {
	// Initialize Database Connection
	ConnectDB()
	defer DB.Close() // Close DB when program exits

	// Register Routes
	RegisterRoutes()

	// Start the server
	fmt.Println("🚀 Server is running on port 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
