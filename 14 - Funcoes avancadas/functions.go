package main

import "fmt"

func calculoMatematico(n1, n2 int)(soma int, subtracao int){
	soma = n1 + n2
	subtracao= n1 - n2
	return soma, subtracao 
}

func main() {
	fmt.Println(calculoMatematico(10,20))

	soma, subtracao:= calculoMatematico(11, 1)
	fmt.Println(soma, subtracao)


}