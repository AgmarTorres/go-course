package main

import "fmt"
type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante  struct{
	pessoa 
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")
  p := pessoa{"Agmar", "Torres", 21, 18}
	fmt.Println(p)

	e := estudante{p, "BCC", "IFSULDEMINAS"}
	fmt.Println(e)
}