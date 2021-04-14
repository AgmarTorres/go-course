package main

import (
	"fmt"
	"time"
)


func main() {
	canal := make( chan string)
	go escrever("Olá Mundo", canal)
	fmt.Println("Depois da função escrever começar a ser executada")
	mensagem := <- canal
	fmt.Println(mensagem)

	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
}

func escrever(texto string, canal chan string) {
	time.Sleep( time.Second * 5 )
	for i := 0; i < 5; i++ {
		canal <- texto
	//	fmt.Printf(texto)
	}

	close(canal)
}