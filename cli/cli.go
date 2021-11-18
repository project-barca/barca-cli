package cli

import (
	"bufio"
	"errors"
	"github.com/urfave/cli/v2"
	"gopkg.in/gookit/color.v1"
	"io"
	"log"
	"os"
)

func GenerateWin() {
	nameProject := os.Args[len(os.Args)-1]

	app := &cli.App{
		Name:        "<seu-projeto>",
		Usage:       "Nomear o nome da pasta de origem para o diretório do projeto",
		Description: "Barca CLI é uma ferramenta para gerar projetos em diversas tecnologias trazendo maior produtividade no desenvolvimento.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang, l",
				Value: "português-brasileiro",
				Usage: "Idioma da exibição",
			},
			&cli.StringFlag{
				Name:  "config, c",
				Usage: "Carregar configurações através de um arquivo `JSON/TOML` ",
			},
		},
		Commands: []*cli.Command{
			//INIT COMMAND *********************  INIT COMMAND   ***************************** INIT COMMAND  ***********************
			{
				Name:        "init",
				Aliases:     []string{"i"},
				Usage:       "Criar projeto",
				Description: "Inicie a tecnologia de seu projeto",
				Subcommands: []*cli.Command{
					//API REST *********************  API REST   ***************************** API REST  ***********************
					{
						Name:        "api",
						Aliases:     []string{"a"},
						Usage:       "Gerar API REST",
						Description: "Implementar Web Server para servir endpoints",
						Action: func(c *cli.Context) error {
							str := c.Args().First()
							color.Yellow.Printf("Gerando API REST...")
							if str == "" {
								return errors.New("Não conseguimos gerar a API REST")
							}
							server, err := os.Open("frameworks/http/gorilla/gorilla-mux.go")
							if err != nil {
								panic(err)
							}
							defer func() {
								if err := server.Close(); err != nil {
									panic(err)
								}
							}()
							r := bufio.NewReader(server)

							_, errorPath := os.Stat("./" + str)

							if os.IsNotExist(errorPath) {
								errDir := os.Mkdir("./"+str, 0777)
								if errDir != nil {
									log.Fatal(errorPath)
									color.Red.Printf("Não houve sucesso ao tentar criar diretório")
									return nil
								}
								currentPath := "./" + str
								color.Green.Printf("Diretório atual: " + currentPath)
							}
							color.Yellow.Printf("Gerando servidor com Gorilla Mux")
							currentServer, err := os.Create(nameProject + "./main.go")
							if err != nil {
								panic(err)
							}
							//defer color.Green.Printf("Projeto criado com sucesso!")
							defer func() {
								if err := currentServer.Close(); err != nil {
									panic(err)
								}
							}()

							w := bufio.NewWriter(currentServer)

							buf := make([]byte, 1024)
							for {
								n, err := r.Read(buf)
								if err != nil && err != io.EOF {
									panic(err)
								}
								if n == 0 {
									break
								}
								if _, err := w.Write(buf[:n]); err != nil {
									panic(err)
								}
							}

							if err = w.Flush(); err != nil {
								panic(err)
							}

							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
