package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var tmpl *template.Template
var db *sql.DB = getMySQLDB()

func getMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@dk@tcp(127.0.0.1:3306)/jdbc?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected")
	return db
}

func init() {
	tmpl = template.Must(template.ParseFiles("crudForm.html"))
}

type studentInfo struct {
	Sid    string
	Name   string
	Course string
}

func crudHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	student := studentInfo{
		Sid:    r.FormValue("Sid"),
		Name:   r.FormValue("name"),
		Course: r.FormValue("course"),
	}

	if r.FormValue("submit") == "insert" {
		Sid, _ := strconv.Atoi(student.Sid)
		_, err := db.Exec("insert into student (Sid,name,course) values (?,?,?)", Sid, student.Name, student.Course)
		if err != nil {
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: err.Error()})
		} else {
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: "Record insert"})
		}

	} else if r.FormValue("submit") == "update" {
		tmpl.Execute(w, struct {
			Success bool
			Message string
		}{Success: true, Message: "Welcome to update"})

	} else if r.FormValue("submit") == "read" {
		data :=[]string{}
		Rows,err := db.Query("select * from student")
		if err != nil{
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{Success: true, Message: err.Error()})
		}else{
			s := studentInfo{}
			data = append(data, "<table border>")
			data = append(data, "<tr></th>Student ID</th></th>Student Name</th></th>Student Course</th><tr>")
			for Rows.Next(){
				Rows.Scan(&s.Sid,&s.Name,&s.Course)
				data = append(data, fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td></tr>", s.Sid, s.Name, s.Course))
			
			}
			tmpl.Execute(w, struct {
				Success bool
				Message string
				Data    string
			}{Success: true, Message: "Welcome to read", Data: strings.Join(data, "")})
		}
		

	} else if r.FormValue("submit") == "delete" {
		tmpl.Execute(w, struct {
			Success bool
			Message string
		}{Success: true, Message: "Welcome to delete"})
	}
	fmt.Println(student)

}

func main() {
	http.HandleFunc("/", crudHandler)
	http.ListenAndServe(":8888", nil)
}
