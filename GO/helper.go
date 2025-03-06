package main
import (
	"fmt"
	"strings"

)

func validateUser(name string,email string ,tickets int){
	isValidName := len(name) >= 2
	isValidEmail := strings.Contains(email,"@")
	isValidTickets := tickets > 0

   if(isValidName && isValidEmail && isValidTickets){
	  fmt.Printf("name : %v , email : %v , ticketes : %v\n",name,email,tickets)
   }else{
	  fmt.Println("InValid credentails")
   }
}