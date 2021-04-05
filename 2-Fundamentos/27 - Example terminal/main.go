//go mod init linha-de-comando
//go get github.com/urfave/cli
package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)



func main(){
	fmt.Println("Ponto de partida")
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}

 