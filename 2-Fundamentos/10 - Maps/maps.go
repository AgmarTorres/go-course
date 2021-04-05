package main

import "fmt"

func main(){
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome": "Agmar",
		"sobrenome":"Silva",

	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome":{
			"primeiro": "João",
			"ultimo":"Pedro",
		},
		"curso":{
			"nome": "Engenharia",
			"campus":"Campus 1",
		},
	}
		
		fmt.Println(usuario2)
		delete(usuario2, "nome")
		fmt.Println(usuario2)

		// Adicionando informação
		usuario2["signo"] = map[string]string{
			"nome": "Gemêos",
		}

		fmt.Println(usuario2)
	
}