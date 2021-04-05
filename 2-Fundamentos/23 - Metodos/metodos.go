package main

import "fmt"

type usuario struct {
	nome      string
	sobrenome string
	idade uint8
}

func ( u usuario) salvar(){
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados \n", u.nome)
}

func ( u usuario) maiorDeIdade() bool{
	return u.idade >= 18
}

func ( u *usuario) fazerAniversario(){
	u.idade++
}

func main() {

	u := usuario{"Agmar", "Torres", 19}
	fmt.Println(u)
	u.salvar()


	maiorDeIdade := u.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	u.fazerAniversario()
	fmt.Println(u)
}