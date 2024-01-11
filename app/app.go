package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Retorna a aplicação de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca de IPs e Nomes de Servidor"

	app.Commands = []cli.Command{
		{
			// Configura o nome do subcomando para "ip"
			Name:  "ip",
			Usage: "Busca por IPs",
			// As flags é como se fosse o parâmetro para o comando funcionar
			Flags: []cli.Flag{
				// Lista de flags associada ao subcomando "ip"
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			// Executada quando o subcomando "ip" é chamado
			Action: buscarIps,
		},
		{
			Name:  "servidor",
			Usage: "Busca por informações do servidor",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: buscarServidor,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	// O método "String" obtém o valor fornecido pela flag --host
	host := c.String("host")

	// Realiza uma pesquisa DNS e obtem os endereços IPs associados ao host
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidor(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) //name server
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
