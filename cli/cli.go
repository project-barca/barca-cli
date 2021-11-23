package cli

import (
	"log"
	"os"

	"github.com/project-barca/barca-cli/generate"
	"github.com/project-barca/barca-cli/generate/node/integrate"
	"github.com/project-barca/barca-cli/generate/node/make"
	"github.com/urfave/cli/v2"
)

func GenerateWin() {
	nameProject := os.Args[len(os.Args)-1]
	var language string
	var framework string
	var port string
	var url string

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
				Name:        "connect, c",
				Usage:       "Especifíque a URL para Conectar Com Serviços",
				Destination: &url,
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
				Usage:       "Criar Um Novo Projeto",
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
			//ADD COMMAND *********************  ADD COMMAND   ***************************** ADD COMMAND  ***********************
			{
				Name:        "add",
				Aliases:     []string{"a"},
				Usage:       "Construir recursos em Projeto",
				Description: "Integrar Serviços para seu Projeto",
				Subcommands: []*cli.Command{
					//MODEL ENTITY *********************  MODEL ENTITY   ***************************** MODEL ENTITY  ***********************
					{
						Name:        "model",
						Aliases:     []string{"m"},
						Usage:       "Adicionar Modelo",
						Description: "Implementar entidades para o projeto",
						Action: func(c *cli.Context) error {
							directory := c.Args().First()
							make.Model(language, directory, framework, port, nameProject)

							return nil
						},
					},
					//CONTROLLER *********************  CONTROLLER   ***************************** CONTROLLER  ***********************
					{
						Name:        "controller",
						Aliases:     []string{"c"},
						Usage:       "Adicionar Controller",
						Description: "Implementar controladores funcionais para o projeto",
						Action: func(c *cli.Context) error {
							directory := c.Args().First()
							make.Controller(language, directory, framework, port, nameProject)

							return nil
						},
					},
				},
			},
			//INTEGRATE COMMAND *********************  INTEGRATE COMMAND   ***************************** INTEGRATE COMMAND  ***********************
			{
				Name:        "integrate",
				Aliases:     []string{"in"},
				Usage:       "Integrar Recursos ao Projeto",
				Description: "Integrar recursos para deixar seu projeto mais robusto",
				Subcommands: []*cli.Command{
					//MySQL *********************  MySQL   ***************************** MySQL  ***********************
					{
						Name:        "mysql",
						Aliases:     []string{"ms"},
						Usage:       "Integrar Banco de Dados MySQL",
						Description: "Implementar MySQL em seu projeto",
						Action: func(c *cli.Context) error {
							directory := c.Args().First()
							integrate.MySQL(language, directory, port, nameProject)

							return nil
						},
					},
					//MongoDB *********************  MongoDB   ***************************** MongoDB  ***********************
					{
						Name:        "mongo",
						Aliases:     []string{"mg"},
						Usage:       "Integrar Banco de Dados MongoDB",
						Description: "Implementar MongoDB NoSQL em seu projeto",
						Action: func(c *cli.Context) error {
							directory := c.Args().First()
							integrate.Mongo(language, directory, framework, port, nameProject)

							return nil
						},
					},
					//PostgreSQL *********************  PostgreSQL   ***************************** PostgreSQL  ***********************
					{
						Name:        "postgresql",
						Aliases:     []string{"pq"},
						Usage:       "Integrar Banco de Dados PostgreSQL",
						Description: "Implementar PostgreSQL em seu projeto",
						Action: func(c *cli.Context) error {
							directory := c.Args().First()
							generate.API(language, directory, framework, port, nameProject)

							return nil
						},
					},
					//Microsoft SQL Server *********************  Microsoft SQL Server   ***************************** Microsoft SQL Server  ***********************
					{
						Name:        "sqlserver",
						Aliases:     []string{"mss"},
						Usage:       "Integrar Banco de Dados Microsoft SQL Server",
						Description: "Implementar Microsoft SQL Server em seu projeto",
						Action: func(c *cli.Context) error {
							directory := c.Args().First()
							generate.API(language, directory, framework, port, nameProject)

							return nil
						},
					},
					//Oracle SQL *********************  Oracle SQL  ***************************** Oracle  SQL ***********************
					{
						Name:        "oracle",
						Aliases:     []string{"oc"},
						Usage:       "Integrar Banco de Dados Oracle SQL",
						Description: "Implementar Oracle SQL em seu projeto",
						Action: func(c *cli.Context) error {
							directory := c.Args().First()
							generate.API(language, directory, framework, port, nameProject)

							return nil
						},
					},
					//AWS DynamoDB *********************  AWS DynamoDB  ***************************** AWS DynamoDB ***********************
					{
						Name:        "dynamo",
						Aliases:     []string{"dy"},
						Usage:       "Integrar Banco de Dados DynamoDB",
						Description: "Implementar Amazon DynamoDB NoSQL em seu projeto",
						Action: func(c *cli.Context) error {
							directory := c.Args().First()
							generate.API(language, directory, framework, port, nameProject)

							return nil
						},
					},
					//IBM DB2 *********************  IBM DB2  ***************************** IBM DB2 ***********************
					{
						Name:        "db2",
						Aliases:     []string{"d2"},
						Usage:       "Integrar Banco de Dados IBM DB2",
						Description: "Implementar IBM DB2 em seu projeto",
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
