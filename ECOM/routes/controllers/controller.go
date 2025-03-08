// package controllers

// import (
// 	"ecom/config"
// 	"ecom/models"
// 	"html/template"
// 	"log"
// 	"net/http"
// )

// func HomePage(w http.ResponseWriter, r *http.Request) {
// 	db := config.ConnectDB()
// 	defer db.Close() // Ensure DB connection is closed

// 	result, err := db.Query(`SELECT * FROM products`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer result.Close() // Close result set

// 	products := []models.Product{}
// 	for result.Next() {
// 		p := models.Product{}
// 		err := result.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		products = append(products, p)
// 	}

// 	tmpl := template.Must(template.ParseGlob("templates/*.html"))
// 	err = tmpl.ExecuteTemplate(w, "index.html", products)
// 	if err != nil {
// 		http.Error(w, "Error rendering template", http.StatusInternalServerError)
// 	}
// }
