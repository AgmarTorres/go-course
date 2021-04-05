package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel2)

	var v3 int  
	var ponteiro *int

	v3 = 10
	ponteiro = &v3

	fmt.Println(v3, ponteiro)
	fmt.Println(v3, &v3)

	v3 = 150
	fmt.Println(v3, *ponteiro)

}