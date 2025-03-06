// package main

// import (
// 	"html/template"
// 	"net/http"
// )

// var temp * template.Template
// func init(){
// 	temp = template.Must(template.ParseGlob("templates/*.html"))
	
// }

// func handleFunc(w http.ResponseWriter,r *http.Request){
// 	// fmt.Fprintf(w,"Welcome to Go Web page")
// 	temp.ExecuteTemplate(w,"WelcomePage.html",nil) // redirecting to templtes page
// }

// func main(){
// 	http.HandleFunc("/",handleFunc)
// 	http.ListenAndServe(":9999",nil)
// }