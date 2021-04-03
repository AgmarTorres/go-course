package main

import "fmt"

func diaSemana( num int) string{
	switch num {
		case 1:
			return "Segunda-feita"
			
		case 2:
			return "Terça-feita"
			
		case 3:
			return "Quarta-feita"
			
		case 4:
			return "Quinta-feita"
			
		case 5:
			return "Sexta-feita"
			
		case 6:
			return "Sabado"	

		default:
			return "Domigno"
	}
}

func main(){
	dia := diaSemana(2)
	fmt.Println(dia)
}