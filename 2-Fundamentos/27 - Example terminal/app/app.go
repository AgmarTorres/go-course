package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)


func Gerar() *cli.App {
		app := cli.NewApp()
		app.Name="Aplicação de linha de comando"
		app.Usage = "Busca por Ips e nomes"
		flags := []cli.Flag{
				cli.StringFlag{
				Name: "host",
				Value: "devbook.com.br",
			},
		}

		app.Commands = []cli.Command{
			{
				Name:"ip",
				Usage : "Busca por Ips e nomes",
				Flags: flags,
				Action: buscaIps,
			},
			{
				Name: "servidores", //go run main.go servidores --host amazon.com.br
				Usage: "Busca o nome dos servidores na internet",
				Flags: flags,
				Action: buscarNomeServidores,
			},
		}
		return app
}

func buscaIps( c *cli.Context){
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil{
		log.Fatal(erro)
	}
	for _, ip:= range ips{
		fmt.Println(ip)
	}
}

func buscarNomeServidores( c *cli.Context){
	host := c.String("host")
	servidores, erro :=net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores{
		fmt.Println(servidor.Host)
	} 
}