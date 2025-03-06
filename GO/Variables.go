// package main
// import (
// 	"fmt"
// )
// const PI float64 = 3.14

// func main(){
// 	variables()
// 	formattedString()
// 	dataTypes()
// 	getUsesInput()
// 	pointers()
// 	array()
// 	slice()
	
// }

// func variables(){
// 	var name string  = "kumar"
// 	age := 19
// 	fmt.Println("Name ",name)
// 	fmt.Println("Age : ",age)
// 	fmt.Println("PI :",PI)
// }

// func formattedString(){
// 	name := "dhinesh"
// 	age := 19
// 	fmt.Printf("Name : %v, Age : %v\n",name,age)
// }

// func dataTypes(){
// 	var x int = 19
// 	var y float64 = 9813.123
// 	var name string = "kumar"
// 	var isValid bool = false
// 	fmt.Printf("x : %v, y : %v , name : %v ,isValid : %v\n",x,y,name,isValid)
// }

// func getUsesInput(){
// 	var name string 
// 	var salary float64
// 	var dept string
// 	fmt.Println("Enter the name")
// 	fmt.Scan(&name)
// 	fmt.Println("Enter the salary")
// 	fmt.Scan(&salary)
// 	fmt.Println("Enter the dept")
// 	fmt.Scan(&dept)

// 	fmt.Printf("Name : %v , Salary : %v , Deptartment : %v\n",name,salary,dept)
// }

// func pointers(){
// 	var x int = 42
// 	var y *int = &x
// 	fmt.Println(x)
// 	fmt.Println(y)
// }

// func array(){
// 	var arr = [3]int{}
// 	fmt.Println("Array values : ",arr)
// } 

// func slice(){
// 	slice := []int{10,20,30}
// 	slice = append(slice, 40)
// 	fmt.Println(slice)
// }