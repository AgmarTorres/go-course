package main

import "fmt"


func main(){
	fmt.Println("Canal com buffer")
	canal := make( chan string,2)
	canal <- "Hello 1"
	canal <- "Hello 2"

	message1 := <- canal
	message2 := <- canal
	fmt.Println(message1)
	fmt.Println(message2)
}