package cli

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/project-barca/barca-cli/generate"
	"github.com/project-barca/barca-cli/generate/node/integrate"
	faca "github.com/project-barca/barca-cli/generate/node/make"
	insert "github.com/project-barca/barca-cli/generate/node/make/add/columns"
	"github.com/project-barca/barca-cli/utils/dir"
	"github.com/project-barca/barca-cli/utils/file"
	"github.com/urfave/cli/v2"
)

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte

	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

func GenerateWin() {
	nameProject := os.Args[len(os.Args)-1]
	var language string
	var framework string
	var database string
	var port string
	var host string
	var url string
	var dbname string
	var user string
	var password string
	var table string
	var field string
	var types string
	var collection string

	app := &cli.App{
		Name:        "barca-cli",
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
				Usage:       "Especifíque a Porta do Servidor",
				Destination: &port,
			},
			&cli.StringFlag{
				Name:        "host, h",
				Value:       "localhost",
				Usage:       "Especifíque a Host do Servidor",
				Destination: &host,
			},
			&cli.StringFlag{
				Name:        "dbname, dn",
				Usage:       "Especifíque o nome do Banco de Dados que deseja utilizar",
				Destination: &dbname,
			},
			&cli.StringFlag{
				Name:        "user, u",
				Usage:       "Especifíque o usuário",
				Destination: &user,
			},
			&cli.StringFlag{
				Name:        "password, p",
				Usage:       "Especifíque a senha do usuário",
				Destination: &password,
			},
			&cli.StringFlag{
				Name:        "database, db",
				Usage:       "Especifíque o Banco De Dados para utilizar no recurso solicitado",
				Destination: &database,
			},
			&cli.StringFlag{
				Name:        "field, fi",
				Usage:       "Defina o nome do campo",
				Destination: &field,
			},
			&cli.StringFlag{
				Name:        "type, t",
				Usage:       "Defina o tipo de dados para o campo",
				Destination: &types,
			},
			&cli.StringFlag{
				Name:        "table, t",
				Usage:       "Especifíque o nome da Tabela para Banco De Dados SQL",
				Destination: &table,
			},
			&cli.StringFlag{
				Name:        "collection, cl",
				Usage:       "Defina o nome da Coleção para Banco De Dados NoSQL",
				Destination: &collection,
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
				Aliases:     []string{"i", "iniciar"},
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
			//INSERT COMMAND *********************  INSERT COMMAND   ***************************** INSERT COMMAND  ***********************
			{
				Name:        "insert",
				Aliases:     []string{"i", "inserir", "acrescentar"},
				Usage:       "Inserir Recurso",
				Description: "Inserir recursos em seu projeto",
				Subcommands: []*cli.Command{
					//COLUMNS MODEL DB *********************  COLUMNS MODEL DB   ***************************** COLUMNS MODEL DB  ***********************
					{
						Name:        "fields",
						Aliases:     []string{"f", "campos", "colunas"},
						Usage:       "Adicionar Novos Campos em Entidade",
						Description: "Adicionar novos atributos em modelo de entidade",
						Action: func(c *cli.Context) error {
							directory := c.Args().First()
							insert.ColumnsMongoModel(directory, collection, field, types)
							return nil
						},
					},
				},
			},
			//ADD COMMAND *********************  ADD COMMAND   ***************************** ADD COMMAND  ***********************
			{
				Name:        "add",
				Aliases:     []string{"a", "adicionar", "acrescentar"},
				Usage:       "Construir recursos em Projeto",
				Description: "Integrar Serviços para seu Projeto",
				Subcommands: []*cli.Command{
					//MODEL ENTITY *********************  MODEL ENTITY   ***************************** MODEL ENTITY  ***********************
					{
						Name:        "model",
						Aliases:     []string{"m", "modelo", "entidade"},
						Usage:       "Adicionar Modelo",
						Description: "Implementar entidades para o projeto",
						Action: func(c *cli.Context) error {
							directory := c.Args().First()
							faca.Model(language, directory, collection, database)

							return nil
						},
					},
					//CONTROLLER *********************  CONTROLLER   ***************************** CONTROLLER  ***********************
					{
						Name:        "controller",
						Aliases:     []string{"c", "controler", "controle"},
						Usage:       "Adicionar Controller",
						Description: "Implementar controladores funcionais para o projeto",
						Action: func(c *cli.Context) error {
							//directory := c.Args().First()
							//faca.Controller(language, directory, framework, port, nameProject)

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
							dbservice := "mysql"
							integrate.MySQL(directory, dbservice, dbname, port, host, user, password, language)

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
							//directory := c.Args().First()
							//integrate.Mongo(language, directory, framework, port, nameProject)

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
							//directory := c.Args().First()
							//generate.API(language, directory, framework, port, nameProject)

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
							//directory := c.Args().First()
							//generate.API(language, directory, framework, port, nameProject)

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
							//directory := c.Args().First()
							//generate.API(language, directory, framework, port, nameProject)

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
							//directory := c.Args().First()
							//generate.API(language, directory, framework, port, nameProject)

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
							//directory := c.Args().First()
							//generate.API(language, directory, framework, port, nameProject)

							return nil
						},
					},
				},
			},
			//SCAN COMMAND *********************  ADD COMMAND   ***************************** ADD COMMAND  ***********************
			{
				Name:        "scan",
				Aliases:     []string{"sc", "scanner", "scanear", "escanear"},
				Usage:       "Escanear Máquina Computacional",
				Description: "Escanear a Máquina Hospedeira dos Projetos e seus Dispositivos Conectados",
				Subcommands: []*cli.Command{
					// SCAN DEVICES *********************  SCAN DEVICES   ***************************** SCAN DEVICES  ***********************
					{
						Name:        "devices",
						Aliases:     []string{"dv", "dispositivos"},
						Usage:       "Escanear Dispositivos",
						Description: "Escanear Dispositivos Conectados",
						Action: func(c *cli.Context) error {
							//directory := c.Args().First()
							fmt.Printf("scan dispositivos conectados")

							return nil
						},
					},
					// SCAN MOTHERBOARD *********************  SCAN MOTHERBOARD   ***************************** SCAN MOTHERBOARD  ***********************
					{
						Name:        "motherboard",
						Aliases:     []string{"mb", "minhaplacamae", "minha-placa", "minhaplaca", "minha-placa-mae", "motherboard", "placa-mae", "placamae", "moba", "baseboard"},
						Usage:       "Escanear Placa-Mãe",
						Description: "Obtenha informações da Placa Mãe e Periféricos Conectados",
						Action: func(c *cli.Context) error {
							//directory := c.Args().First()

							cmd := exec.Command("dir", "C:\\", ">", "D:\\info.txt")
							//cmd := exec.Command("ipconfig", "/all")
							//cmd := exec.Command("wmic", "cpu", "get", "name")

							//FABRICANTE DA PLACA MAE
							//cmd := exec.Command("wmic", "baseboard", "get", "manufacturer")

							//NOME DO PRODUTO DA PLACA
							//cmd := exec.Command("wmic", "baseboard", "get", "product")

							//VERSÃO DA PLACA
							//cmd := exec.Command("wmic", "baseboard", "get", "version")

							//NÚMERO DE SÉRIA DA PLACA
							//cmd := exec.Command("wmic", "baseboard", "get", "serialnumber")

							//VERSAO DA BIOS
							//cmd := exec.Command("wmic", "bios", "get", "smbiosbiosversion")

							var stdout, stderr []byte
							var errStdout, errStderr error
							stdoutIn, _ := cmd.StdoutPipe()
							stderrIn, _ := cmd.StderrPipe()
							err := cmd.Start()
							if err != nil {
								log.Fatalf("cmd.Start() failed with '%s'\n", err)
							}

							// cmd.Wait() should be called only after we finish reading
							// from stdoutIn and stderrIn.
							// wg ensures that we finish
							var wg sync.WaitGroup
							wg.Add(1)
							go func() {
								stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
								wg.Done()
							}()

							stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)

							wg.Wait()

							err = cmd.Wait()
							if err != nil {
								log.Fatalf("cmd.Run() failed with %s\n", err)
							}
							if errStdout != nil || errStderr != nil {
								log.Fatal("failed to capture stdout or stderr\n")
							}
							outStr, errStr := string(stdout), string(stderr)
							fmt.Printf("\n%s\nerr:\n%s\n", outStr, errStr)

							return nil
						},
					},
					// SCAN NETWORK *********************  SCAN NETWORK   ***************************** SCAN NETWORK  ***********************
					{
						Name:        "network",
						Aliases:     []string{"net", "rede", "networks", "mynet", "my-net", "minhanet", "minha-rede", "minharede", "interfaces-de-rede", "interfacesderede", "interface-de-rede", "interfacederede"},
						Usage:       "Escanear Rede",
						Description: "Obtenha Informações Sobre a Infraestrutura de Redes Conectada",
						Action: func(c *cli.Context) error {
							//directory := c.Args().First()
							//	faca.Model(language, directory, collection, database)
							fmt.Printf("scan net")
							return nil
						},
					},
					// SCAN PROJETO *********************  SCAN PROJETO   ***************************** SCAN PROJETO  ***********************
					{
						Name:        "project",
						Aliases:     []string{"proj", "projeto", "projet", "meuprojeto", "meu-projeto", "project-me", "meu-aplicativo", "meuaplicativo", "minhaapi", "minha-api", "me-api", "minhaaplicacao", "minhaaplicacap", "minha-aplicacao", "minha-aplicacap"},
						Usage:       "Escanear Projeto",
						Description: "Obtenha Informações Sobre Tecnologias, Vulnerabilidades, Atualizações, Ramificações pendentes em seu Projeto",
						Action: func(c *cli.Context) error {
							project := c.Args().First()

							if dir.IsDirEmpty(project) == true {
								fmt.Print("O caminho do Projeto específicado está vazio")
							} else {
								fmt.Print("Este projeto contém arquivos e pastas")
								file.NewTempFile("barca-cli.dir-"+project+".txt", "BARCA_TEMP_LIST_DIR")
								dir.Scanner(project)

								input, err := ioutil.ReadFile(os.Getenv("BARCA_TEMP_LIST_DIR"))
								if err != nil {
									log.Println(err)
								}
								lines := strings.Split(string(input), "\n")

								var files []string
								for i := 0; i < len(lines); i++ {
									files = append(files, filepath.Ext(lines[i]))
								}
								fmt.Print(file.WhatIsExtension(files))
								//file.GetInfo(project + "/express.js")
							}

							return nil
						},
					},
					// SCAN FOLDER *********************  SCAN FOLDER   ***************************** SCAN FOLDER  ***********************
					{
						Name:        "folder",
						Aliases:     []string{"fol", "pasta", "caminho", "folder"},
						Usage:       "Escanear Pasta",
						Description: "Obtenha Informações Sobre Tamanho de Arquivos, Permissões, ",
						Action: func(c *cli.Context) error {
							folder := c.Args().First()
							if dir.IsDirEmpty(folder) == true {
								fmt.Print("O caminho do Projeto específicado está vazio")
							} else {
								fmt.Print("Este projeto contém arquivos e pastas")
								file.NewTempFile("barca-cli.dir-"+folder+".txt", "BARCA_TEMP_LIST_DIR")
								dir.Scanner(folder)

								input, err := ioutil.ReadFile(os.Getenv("BARCA_TEMP_LIST_DIR"))
								if err != nil {
									log.Println(err)
								}
								lines := strings.Split(string(input), "\n")

								var files []string
								for i := 0; i < len(lines); i++ {
									files = append(files, filepath.Ext(lines[i]))
								}
								fmt.Print(file.WhatIsExtension(files))
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
