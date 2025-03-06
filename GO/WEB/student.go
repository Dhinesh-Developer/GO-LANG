// package main

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// )

// // Load the template once
// var tmpl *template.Template

// func init() {
// 	tmpl=template.Must(template.ParseGlob("templates/*.html"))
	
// }

// // StudentInfo struct with exported fields (uppercase)
// type StudentInfo struct {
// 	ID     string
// 	Name   string
// 	Course string
// }

// func studentHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		// Serve the form for GET requests
// 		err := tmpl.Execute(w, nil)
// 		if err != nil {
// 			http.Error(w, "Error rendering template", http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		// Read form data
// 		student := StudentInfo{
// 			ID:     r.FormValue("id"),
// 			Name:   r.FormValue("name"),
// 			Course: r.FormValue("course"),
// 		}

// 		// Send data back to template
// 		err := tmpl.Execute(w, struct {
// 			Success bool
// 			Student StudentInfo
// 		}{true, student})

// 		if err != nil {
// 			http.Error(w, "Error rendering template", http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	// If method is not GET or POST
// 	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// }

// func main() {
// 	http.HandleFunc("/", studentHandler)
// 	log.Println("Server running on http://localhost:8888")
// 	err := http.ListenAndServe(":8008", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
