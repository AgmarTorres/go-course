package main

import (
	"fmt"
	"reflect"
)


func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posicação 1","Posicação 2", "Posicação 3", "Posicação 4", "Posicação 5" }
	fmt.Println(array2)

		array3 := [...]int{1,2,3,4,5}
		fmt.Println(array3)

		
	slice := []int{10,11,12,13,14,15}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)


	slice2 := array2[1:3]
	fmt.Println(slice2)

	slice2 [ 1] = "Posição alterada"
	fmt.Println( slice2)
	fmt.Println(array2)

	//Arrays Internos
	//Função make armazena um espaço na memória

	slice3 := make([]float32, 10, 15) // Array float com 10 posições até 15
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4= append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}