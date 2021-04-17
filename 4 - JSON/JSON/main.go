package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)


type cachorro struct {
	Nome     string `json:"nome"`
	Rome     string `json:"raca"`
	Idadeome uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 4}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJSON)
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Puddle",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))

	
}