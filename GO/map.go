package main

import(
	"fmt"
)

func mapExample(){
	studentMarks := make(map[string]int)

	studentMarks["Alice"]= 85
	studentMarks["bob"] = 75
	studentMarks["charlie"] = 90

	fmt.Println("Marks of bob: ",studentMarks["Bob"])
	delete(studentMarks,"charlie")

	fmt.Println("Students marks")
	for name,mark := range studentMarks{
		fmt.Println(name,": ",mark)
	}
}