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
		app.Commands = []cli.Command{
			{
					Name: "ip", //go run main.go ip --host amazon.com.br
					Usage: "Busca Ips de endereços na internet",
					Flags: []cli.Flag{
						cli.StringFlag{
						Name: "host",
						Value: "devbook.com.br",
					},
				},
					Action: buscaIps, 
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