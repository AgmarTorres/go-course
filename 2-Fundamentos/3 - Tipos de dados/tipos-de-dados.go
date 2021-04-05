package main

import "fmt"

func main() {
	var numero int16 = -100
	fmt.Println(numero)
	// int8, int16, int32, int64
	//unsygned int
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	var numero4  byte = 123
	fmt.Println(numero4)

	//float32, float64

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 1234.67
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	str2:= "Texto2"
	fmt.Println(str2)

	char:= '%'
	fmt.Println(char)

	// FIM STRINGS

	//var texto string
	//fmt.Println(texto)

	texto := 5
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error 
	fmt.Println(erro)
}