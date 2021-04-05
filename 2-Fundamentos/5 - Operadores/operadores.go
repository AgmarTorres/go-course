package main

import "fmt"

func main() {
	// Aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 3 % 3
	fmt.Println(soma, subtracao,divisao,  multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25
	
	sum := numero1 + int16(numero2)
	fmt.Println(sum)

	// Atribuição

	var variavel1  string = "String"
	variavel2 := "String 2"
	fmt.Println(variavel1, variavel2)

//Operadores relacionais

	fmt.Println(1>2)
	fmt.Println(1<2)
	fmt.Println(1==2)
	fmt.Println(1!=2)
	fmt.Println(1>=2)
	fmt.Println(1<=2)

	// Operadores Lógicos

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	
	//Operadores unários

	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 15
	fmt.Println(numero)

	numero++
	fmt.Println(numero)
	numero *= 15
	fmt.Println(numero)

	// Fim dos Operadores unários

	//texto := numero > 5 ? "Maior que 5" : "Menor que 5"   não existe
	var text string
	if numero > 5 {
		text = "Maior que 5"
	}	else{
		text = "Menor que 5"
	}
	fmt.Println(text)
}