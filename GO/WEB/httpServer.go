// package main

// import (
// 	"fmt"
// 	"net/http"
// )


// func handlerFunc(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w,"Welcome to First web page in goLang %s",r.URL.Path)
// }

// func main() {
// 	http.HandleFunc("/",handlerFunc)
// 	http.ListenAndServe(":8080",nil)
// }