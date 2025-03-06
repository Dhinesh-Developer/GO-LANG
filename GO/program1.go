// package main

// import (
// 	"fmt"
// 	"os/user"
// )

// func main(){

// 	mapExample()
// 	// variable
// 	// var conferenceName = "Go conference"
// 	// const conferenceTickets = 50
// 	// var remainingTickets = 50

// 	// // Finding the types of the variable
// 	// fmt.Printf("conferenceName %T, conferenceTickets %T, remainingTickets %T\n",conferenceName,conferenceTickets,remainingTickets)
// 	// fmt.Println("Welcome to our conference",conferenceName," booking applications")
// 	// fmt.Printf("We have total of %v tickets and %v tickets are available\n",conferenceTickets,remainingTickets);
// 	// fmt.Println("Get Your tickets to attend")

// 	// //Data types.
// 	// var userName string
// 	// var userTickets int
// 	// var firstName string
// 	// var lastName string
// 	// //ask user for the name

// 	// userName = "Tom"
// 	// userTickets = 2
// 	// fmt.Println(userName)
// 	// fmt.Println(userTickets)

// 	// // getting th user input
	 
// 	//  fmt.Println(userName)
// 	//  fmt.Println(&userName) // ->pointer the memoery location
// 	//  fmt.Println("Enter the user Name")
// 	//  fmt.Scan(&userName)  // ->Getting the user input.
	
// 	//  fmt.Printf("User %v booked %v tickets.\n",userName,userTickets)

// 	//  // book tickets logic
// 	//  remainingTickets = remainingTickets - userTickets;
// 	//  fmt.Println("Thank you ...",remainingTickets)

// 	//  // Arrays ans slices
// 	// //  var bookings = [50]string{"Name","Niclose","Peter"}
	 
// 	//  var bookings [50]string
	
// 	//  fmt.Println("Enter the first Name")
// 	//  fmt.Scan(&firstName)
// 	//  fmt.Println("Enter the last Name")
// 	//  fmt.Scan(&lastName)
// 	//  bookings[0] = firstName+" "+lastName

// 	//  fmt.Printf("The whole array : %v\n", bookings)
// 	//  fmt.Printf("The first name value : %v\n",bookings[0])
// 	//  fmt.Printf("Array type : %T", bookings)
// 	//  fmt.Printf("Array length : %v",len(bookings))

// 	//  // Slicing -> for dynamic array.
// 	//  var book[]string
// 	//  book = append(book,firstName+" "+lastName)
// 	//  fmt.Println("Slicing")

// 	 // loops in the go. only for loop is allowed
// 	//  for {
		
// 	//  }

// 	// // if else
// 	// var data int =5;
// 	// if(data > 0){
// 	// 	fmt.Println("Positive number")
// 	// }else{
// 	// 	fmt.Println("Negative number")
// 	// }

// 	// d:="Dhinesh";
// 	// if(d == "Dhinesh"){
// 	// 	fmt.Println("equal")
// 	// }else{
// 	// 	fmt.Println("not equal")
// 	// }

// 	//  var da bool = false;
// 	//  if(da!=false){
// 	// 	fmt.Println("boolean is true")
// 	//  }else{
// 	// 	fmt.Println("boolean is false")
// 	//  }

// 	//  // user input validataion
// 	//  greetUsers()
// 	//  var name string 
// 	//  var email string
// 	//  var tickets int
// 	//  getUserInput(name,email,tickets)

// 	//  validateUser(name,email,tickets)

// 	//  // switch case
// 	//  city := "London"

// 	//  switch city{
// 	//  	case "New york":
// 	// 		fmt.Println("New York case 1")
// 	// 		break
// 	// 	case "London":
// 	// 		fmt.Println("London case2")
// 	// 		break
// 	// 	default:
// 	// 		fmt.Println("No valid city selected")	
// 	//  }

// 	//  conferenceName := "go conference"
// 	//  fmt.Println("Enter the Conference Name")
// 	//  fmt.Scan(&conferenceName)
// 	// //  greetUser(confereneceName)

// 	//  var res string = greet()
// 	//  fmt.Println(res)




// }
// // Functions
// func greetUsers(){
// 	fmt.Println("Welcome to our conference")
// }

// // Functions parameter
// func greetUser(conferenceName string){
// 	fmt.Println("Welcome to ",conferenceName)
// }

// // returing th function
// func greet()(string){
// 	return "Go conference"
// }

