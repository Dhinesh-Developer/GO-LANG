package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tmpl * template.Template
func init(){
	tmpl=template.Must(template.ParseGlob("templates/*.html"))
}

type studentInfo struct{
	Sid string
	Name string
	Course string
}

func crudHandler(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost{
		tmpl.Execute(w,nil)
		return 
	}

	student := studentInfo{
		Sid : r.FormValue("Sid"),
		Name : r.FormValue("Name"),
		Course : r.FormValue("Course"),
	}

	fmt.Println(student)

	tmpl.Execute(w,struct{
		Sucess bool
		Message string
	}{Sucess: true,Message: "Welcome"})
}

func main() {
	http.HandleFunc("/stud",crudHandler)
	http.ListenAndServe(":1111",nil)
}