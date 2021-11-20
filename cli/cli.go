package cli

import (
	"log"
	"os"

	"github.com/project-barca/barca-cli/generate"
	"github.com/urfave/cli/v2"
)

func GenerateWin() {
	nameProject := os.Args[len(os.Args)-1]
	var language string
	var framework string
	var port string

	app := &cli.App{
		Name:        "<seu-projeto>",
		Usage:       "Crie uma nova pasta de origem para o diretório do projeto",
		Description: "Barca CLI é uma ferramenta para gerar projetos em diversas tecnologias trazendo maior produtividade no desenvolvimento.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang, l",
				Value:       "portugues-brasileiro",
				Usage:       "Idioma da exibição",
				Destination: &language,
			},
			&cli.StringFlag{
				Name:        "framework, f",
				Value:       "barca",
				Usage:       "Especifíque Framework para servir sua API",
				Destination: &framework,
			},
			&cli.StringFlag{
				Name:        "port, p",
				Value:       "4200",
				Usage:       "Especifíque a Porta do Servidor",
				Destination: &port,
			},
			&cli.StringFlag{
				Name:  "config, c",
				Usage: "Carregar configurações através de arquivo `JSON/TOML` ",
			},
		},
		Commands: []*cli.Command{
			//INIT COMMAND *********************  INIT COMMAND   ***************************** INIT COMMAND  ***********************
			{
				Name:        "init",
				Aliases:     []string{"i"},
				Usage:       "Criar um projeto de software",
				Description: "Inicie a tecnologia que você deseja",
				Subcommands: []*cli.Command{
					//API REST *********************  API REST   ***************************** API REST  ***********************
					{
						Name:        "api",
						Aliases:     []string{"a"},
						Usage:       "Gerar API REST",
						Description: "Implementar Web Server para servir endpoints",
						Action: func(c *cli.Context) error {
							directory := c.Args().First()
							generate.API(language, directory, framework, port, nameProject)

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
