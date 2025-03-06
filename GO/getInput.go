package main
import (
	"fmt"

)

func getUserInput(name string,email string,tickets int){
	
	 fmt.Println("Enter the name")
	 fmt.Scan(&name)
	 fmt.Println("Enter the email")
	 fmt.Scan(&email)
	 fmt.Println("Enter the tickets")
	 fmt.Scan(&tickets)
}

func sendTickets(tickets int,name string){
	fmt.Printf("%v tickets for %v",tickets,name)
}