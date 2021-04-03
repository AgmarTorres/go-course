package main

import "fmt"

func aluEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6!")

}

func recuperarExecucao(){
	fmt.Println("Tentativa de recuperar")
	
	if r := recover(); r != nil {

	}
}

func main() {
	fmt.Println(aluEstaAprovado(3,7))
	fmt.Println("Pós teste")
	fmt.Println(aluEstaAprovado(6,6))
	fmt.Println("Pós teste")

}