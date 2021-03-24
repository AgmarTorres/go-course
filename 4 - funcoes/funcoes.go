package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8,int8,int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	div := n1 / n2
	mult := n1 * n2
	return soma, sub, div, mult
}

func main(){
	fmt.Println( somar(3, 5) )
	sum := somar(10, 20)
	fmt.Println(sum)
	var text string  ="Ola"
	

	var f = func( txt string){
		fmt.Println(txt)
	}
	f("Texto 1")

	resultSoma, resultSub, resultDiv, resultMult := calculos(15, 10)
	fmt.Println(resultSoma, resultSub, resultDiv, resultMult)

	f("Texto da função 1")
	fmt.Println(text)
}