package cli

import (
	"bufio"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"gopkg.in/gookit/color.v1"
	"io"
	"log"
	"os"
	"time"
)

func GenerateWin() {
	nameProject := os.Args[len(os.Args)-1]

	// logsArray := [3]string{"Verificando requisitos...", "Instalando Node.js v14.0.3", "Finalizando"}

	// for index, element := range logsArray {
	// 	fmt.Println(index)
	// 	fmt.Println(element)
	// 	time.Sleep(time.Millisecond * 250)
	// }

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
							spinnerSuccess, _ := pterm.DefaultSpinner.Start("Iniciando Gorilla Mux API...")
							time.Sleep(time.Second * 2)
							spinnerSuccess.Success()

							if str == "" {
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

								currentServer, err := os.Create(nameProject + "./main.go")

								spinnerSuccessDir, _ := pterm.DefaultSpinner.Start("Criando estrutura de pastas...")
								time.Sleep(time.Second * 2)
								spinnerSuccessDir.UpdateText("Projeto: " + nameProject)
								time.Sleep(time.Second * 1)
								spinnerSuccessDir.Success("Projeto construído")
								time.Sleep(time.Second * 1)
								spinnerSuccessDir.Success("OK")

								if err != nil {
									panic(err)
								}
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
							} else {
								pterm.Error.Println("Já existe uma pasta com o mesmo nome do projeto que você está tentando criar")

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
