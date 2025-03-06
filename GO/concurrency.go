// package main
// import (
// 	"fmt"
// 	"time"
// )

// func main(){
// 	go concurrecny1()
// 	time.Sleep(time.Second)


// 	for i:=1;i<=5;i++{
// 		go primenumber(i)
// 	}
// 	time.Sleep(time.Second)
// }

// func concurrecny1(){
// 	fmt.Println("Hello from Go routines")
// }

// func primenumber(i int){
// 	fmt.Println(i)
// }