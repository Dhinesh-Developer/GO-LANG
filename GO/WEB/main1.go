// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// 	// "github.com/gorilla/mux"
// )

// // mux -> package is used to creating the routing in the GO language 

// var tmpl *template.Template

// func init(){
// 	tmpl=template.Must(template.ParseGlob("templates/*.html"))
// }

// func homeHandler(w http.ResponseWriter,r *http.Request){
// 	tmpl.ExecuteTemplate(w,"index.html",nil)
// }

// func contactHandler(w http.ResponseWriter,r *http.Request){
// 	tmpl.ExecuteTemplate(w,"contact.html",nil)
// }

// func aboutHandler(w http.ResponseWriter,r *http.Request){
// 	tmpl.ExecuteTemplate(w,"About.html",nil)
// }

// func handleFunc(w http.ResponseWriter,r *http.Request){
// 	fmt.Fprintf(w,"This is using mux of github")
// }

// func main() {

// 	//Acessing the static folders using the http.Dir()
// 	fs:=http.FileServer(http.Dir("assessts"))
// 	http.Handle("/assessts/",http.StripPrefix("/assessts/",fs))

// 	// mux package for creating a restful api routing

// 	// r := mux.NewRouter()
// 	// r.HandleFunc("/",handleFunc)
// 	// http.ListenAndServe(":8888",r)
// 	// r.PathPrefix("/assessts").Handler(http.StripPrefix("/assessts/",fs)) -> for accessing the static folders using the mux routing instead using the http default router

// 	http.HandleFunc("/home",homeHandler)
// 	http.HandleFunc("/contact",contactHandler)
// 	http.HandleFunc("/about",aboutHandler)
// 	http.ListenAndServe(":8888",nil)
// }

// // go run main.go
// // go run .\withmux.go
// // go get -u github.com/gorilla/mux