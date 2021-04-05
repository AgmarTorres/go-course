package main

import "fmt"

func soma(numeros ...int) int{
	fmt.Println(numeros)
	total := 0	
	for _, numero := range numeros{
			total += numero
		}
	return total
}

func escrever( texto string, numeros ...int)string{
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return fmt.Sprintf("%d%s", total, texto)
}

func main() {
	fmt.Println(soma(1,2,3,4,5,6)	)
	value := escrever("Olá", 1,2,3,4,5,6)
	fmt.Println(value)

}