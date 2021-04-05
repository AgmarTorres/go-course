package main

import "fmt"

type usuario struct{
	nome string
	idade uint8
	endereco endereco
}

type  endereco struct{
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")
  exEndereco := endereco{"Rua X", 80}
	var user usuario
	user.nome="Agmar"
	user.idade = 15
	fmt.Println(user)

	u2 := usuario{"Maria", 21, exEndereco}
	fmt.Println(u2)

	usuario3 := usuario{idade : 21}
	fmt.Println(usuario3)
}