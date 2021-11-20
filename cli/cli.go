package cli

import (
	"log"
	"os"

	"barca-cli/generate"

	"github.com/urfave/cli/v2"
)

func GenerateWin() {
	nameProject := os.Args[len(os.Args)-1]
	var language string

	app := &cli.App{
		Name:        "<seu-projeto>",
		Usage:       "Nomear o nome da pasta de origem para o diretório do projeto",
		Description: "Barca CLI é uma ferramenta para gerar projetos em diversas tecnologias trazendo maior produtividade no desenvolvimento.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang, l",
				Value:       "português-brasileiro",
				Usage:       "Idioma da exibição",
				Destination: &language,
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
							directory := c.Args().First()

							generate.API(language, directory, nameProject)

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
