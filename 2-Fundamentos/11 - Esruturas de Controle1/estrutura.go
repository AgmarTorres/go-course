package main

import "fmt"

func main(){

 numero:= 10

 if numero > 15 {
	 fmt.Println("Maior ")
 } else{
	 fmt.Println("Menor ")
 }

 if outroNumero := numero; outroNumero > 0 {
	 fmt.Println( "Numero é maior que zero")
 }else if numero < 10 {
	 fmt.Println("Número é menor que zero")
 }else{
	 fmt.Println("Entre 0 e -10")
 }

}