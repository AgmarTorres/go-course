package main

import "fmt"

func main() {
	var variavel string = "Variável 1"
	variavel2:="Variável 2" // declaração implicita
	fmt.Println(variavel)
	fmt.Println(variavel2)

	var (
		variavel3 string ="v3"
		variavel4 string = "v4"
	)
	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "v5", "v6"
	fmt.Println(variavel5, variavel6)

	const constante string = "Contant 1"
	fmt.Println(constante)
}