package integrate

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	scan "github.com/project-barca/barca-cli/scan/environment"
	"github.com/project-barca/barca-cli/utils/file"
	projeto "github.com/project-barca/barca-cli/utils/project"
	"github.com/pterm/pterm"

	"golang.org/x/text/language"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func Docker(directory, languageProgram, port, host, user, password, lang string, languagesProgram []string) {
	// Configurações para internacionalização do software
	var hasNode = scan.EnvironmentNode(directory)
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("translate/en.json")
	bundle.LoadMessageFile("translate/fr.json")
	bundle.LoadMessageFile("translate/pt-BR.json")

	switch lang {
	case "portugues-brasileiro":
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	case "english":
		localizer = i18n.NewLocalizer(bundle, language.English.String(), language.BrazilianPortuguese.String(), language.French.String())
	case "francais":
		localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String(), language.BrazilianPortuguese.String())
	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

	switch languageProgram {
	case "python":
		py := [9]string{"flask", "django", "bottle", "falcon", "cherrypy", "pyramid", "web2py", "tornado", "cubicweb"}
		for _, e := range py {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				path := directory
				switch e {
				case "flask":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM python:3.9-alpine3.14\n\nCOPY . /var/www/" + directory + "\n\nRUN apk add --no-cache --virtual .build-deps gcc libc-dev make \\\n  && pip install --no-cache-dir -r /var/www/" + directory + "/requirements.txt \\\n && apk del .build-deps gcc libc-dev make\n\nENTRYPOINT flask run\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "EXPOSE 5000"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)

					if err3 != nil {
						log.Fatal(err3)
					}
				case "django":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM python:3.9-alpine3.14\n\nENV PYTHONDONTWRITEBYTECODE=1\nENV PYTHONUNBUFFERED=1\n\nCOPY . /var/www/" + directory + "\n\nRUN apk add --no-cache --virtual .build-deps gcc libc-dev make \\\n  && pip install --no-cache-dir -r /var/www/" + directory + "/requirements.txt \\\n && apk del .build-deps gcc libc-dev make\n\nENTRYPOINT python manage.py runserver 0.0.0.0:8000 run\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "EXPOSE 5000"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)

					if err3 != nil {
						log.Fatal(err3)
					}
				case "bottle":
					fmt.Print(e)
				default:
					os.Exit(0)
				}
				fmt.Print(e)
			default:
				os.Exit(0)
			}
		}
	case "javascript":
		js := [14]string{"express", "react", "angular", "vue", "ember", "@adonisjs/framework", "rapi", "koa", "meteor", "feather", "sails", "nest", "socket", "derby"}
		for _, e := range js {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				path := directory
				switch e {
				case "adonis":
					fmt.Print(e)
				case "rapi":
					fmt.Print(e)
				case "express":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM node:latest\n\nCOPY . /var/www/" + directory + "\n\nENV PORT=3000\n\nRUN npm install\n\nEXPOSE $PORT\n\nENTRYPOINT [\"node\", \"express.js\"]\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "EXPOSE 5000"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)

					if err3 != nil {
						log.Fatal(err3)
					}
				default:
					os.Exit(0)
				}
			default:
				os.Exit(0)
			}
		}
	case "php":
		php := [8]string{"laravel/framework", "symfony/symfony", "slim/slim", "zend", "phalcon", "cakephp/cakephp", "yii", "codeigniter"}
		for _, e := range php {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				fmt.Print(e)
			default:
				os.Exit(0)
			}
		}
	case "java":
		jv := [8]string{"spring", "play", "gwt", "hibernate", "struts", "vaadin", "yii", "codeigniter"}
		for _, e := range jv {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				fmt.Print(e)
			default:
				os.Exit(0)
			}
		}
	case "go":
		golang := [8]string{"gorilla", "echo", "gin", "beego", "fasthttp", "kit", "fiber", "iris"}
		for _, e := range golang {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				fmt.Print(e)
			default:
				os.Exit(0)
			}
		}
	case "elixir":
		golang := [9]string{"phoenix", "nerves", "sugar", "hedwig", "plug", "trot", "placid", "maru", "kitto"}
		for _, e := range golang {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				fmt.Print(e)
			default:
				os.Exit(0)
			}
		}
	case "ruby":
		rb := [10]string{"rubyonrails", "sinatra-activerecord", "roda", "camping", "cuba", "nancy", "remaze", "goliath", "hanami", "padrinho"}
		for _, e := range rb {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				fmt.Print(e)
			default:
				os.Exit(0)
			}
		}
	case "kotlin":
		rb := [8]string{"ktor", "kweb", "javalin", "spark", "spring", "jooby", "hexagon", "tekniq"}
		for _, e := range rb {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				fmt.Print(e)
			default:
				os.Exit(0)
			}
		}
	case "c":
		fmt.Print("Frame C")
	case "cplusplus":
		fmt.Print("Frame C++")
	case "c#":
		fmt.Print("Frame C#")
	case "rust":
		fmt.Print("Frame Rust")
	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

	switch lang {
	case "portugues-brasileiro":
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	case "english":
		localizer = i18n.NewLocalizer(bundle, language.English.String(), language.BrazilianPortuguese.String(), language.French.String())
	case "francais":
		localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String(), language.BrazilianPortuguese.String())
	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

	_, errorPath := os.Stat("./" + directory)

	localizeConfigErrorMySQLExpress := i18n.LocalizeConfig{
		MessageID: "error_mysql_express",
	}
	localizeConfigSuccessMySQLExpress := i18n.LocalizeConfig{
		MessageID: "success_mysql_express",
	}
	resultErrorMySQLExpress, _ := localizer.Localize(&localizeConfigErrorMySQLExpress)
	resultSuccessMySQLExpress, _ := localizer.Localize(&localizeConfigSuccessMySQLExpress)

	if hasNode == true {
		indexLineDependencies := file.FindPositionLineText(directory+"/package.json", "dependencies")
		file.InsertStringToFile(directory+"/package.json", "    \"mysql2\": \"^2.0.2\",\n    \"sequelize\": \"^5.21.2\",\n", indexLineDependencies)

		if os.IsNotExist(errorPath) {
			pterm.Error.Println(resultErrorMySQLExpress)
		} else {
			path := directory + "/config"

			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Mkdir(path, 0755)
			}
			f, err := os.Create(path + "/Dockerfile")

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			firstLine := "module.exports = {\n  HOST:  '" + host + "',\n  USER:  '" + user + "',\n  PASSWORD:  '" + password + "',\n  DB:  '\n"
			data := []byte(firstLine)

			_, err2 := f.Write(data)

			if err2 != nil {
				log.Fatal(err2)
			}

			line2 := "};"
			data2 := []byte(line2)

			var idx int64 = int64(len(data))

			_, err3 := f.WriteAt(data2, idx)

			if err3 != nil {
				log.Fatal(err3)
			}

			pterm.Success.Println(resultSuccessMySQLExpress)
		}
	} else {
		fmt.Print("NO Runtime Node")
		fmt.Print("NO Runtime Deno")
		fmt.Print("NO Runtime JDK")
		fmt.Print("NO Runtime GO...")
	}

}
