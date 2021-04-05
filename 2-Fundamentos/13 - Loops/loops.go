package main

import (
	"fmt"
)

func main(){
	//i := 0
	//for i < 10 {
	//	fmt.Println(i)
	//	i++
	//	time.Sleep(time.Second)
	//}

//	for j := 0; j < 10; j++ {
//		fmt.Println("Incrementando j", j)
//		time.Sleep(time.Second)
//	}

		nomes := [3]string{"João", "Agmar", "Davi"}

		for  _, nome:= range nomes{
			fmt.Println(nome)
		}

		for indice, letra := range "PALAVRA"{
			fmt.Println(indice, letra)
		}

		usuario := map[ string ] string{
			"nome": "Agmar",
			"sobrenome": "Torres",
		}

		for chave,valor := range usuario {
			fmt.Println(chave, valor)
		}

		type usuarioStruct struct {
			nome string
			sobrenome string

		}

		usuario2 := usuarioStruct{"Agmar", "Junior"}

		// for chave, valor := range usuario2{
		// 	fmt.Println(chave, valor)
		// }
}