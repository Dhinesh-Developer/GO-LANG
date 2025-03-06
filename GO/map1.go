package main
import "fmt"

type Employee struct{
	ID int
	Salary float64
}

func employee(){
	employees := map[string]Employee{
		"John" : {ID : 101, Salary :987921},
		"Bob" : {ID :102 ,Salary: 901234},
		"Charlie" : {ID :102,Salary: 9871324},
	}

	fmt.Println("Employee Details")

	for name,emp := range employees{
		fmt.Printf("Name :%v ,ID :%v , Salary :%v",name,emp.ID,emp.Salary)
	}
}