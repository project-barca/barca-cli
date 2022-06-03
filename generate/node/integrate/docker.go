package integrate

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	projeto "github.com/project-barca/barca-cli/utils/project"

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
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM python:3.9-alpine3.14\n\nCOPY . /var/www/" + directory + "\n\nRUN apk add --no-cache --virtual .build-deps gcc libc-dev make \\\n  && pip install --no-cache-dir -r /var/www/" + directory + "/requirements.txt \\\n && apk del .build-deps gcc libc-dev make\n\nENTRYPOINT python main.py\n\n"
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
				case "pyramid":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM python:3.9-alpine3.14\n\nCOPY . /var/www/" + directory + "\n\nRUN apk add --no-cache --virtual .build-deps gcc libc-dev make \\\n  && pip install --no-cache-dir -r /var/www/" + directory + "/requirements.txt \\\n && apk del .build-deps gcc libc-dev make\n\nENTRYPOINT python main.py\n\n"
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
				case "tornado":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM python:3.9-alpine3.14\n\nCOPY . /var/www/" + directory + "\n\nRUN apk add --no-cache --virtual .build-deps gcc libc-dev make \\\n  && pip install --no-cache-dir -r /var/www/" + directory + "/requirements.txt \\\n && apk del .build-deps gcc libc-dev make\n\nENTRYPOINT python main.py\n\n"
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
				case "cherrypy":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM python:3.9-alpine3.14\n\nCOPY . /var/www/" + directory + "\n\nRUN apk add --no-cache --virtual .build-deps gcc libc-dev make \\\n  && pip install --no-cache-dir -r /var/www/" + directory + "/requirements.txt \\\n && apk del .build-deps gcc libc-dev make\n\nENTRYPOINT python main.py\n\n"
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
				case "falcon":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM python:3.9-alpine3.14\n\nCOPY . /var/www/" + directory + "\n\nRUN apk add --no-cache --virtual .build-deps gcc libc-dev make \\\n  && pip install --no-cache-dir -r /var/www/" + directory + "/requirements.txt \\\n && apk del .build-deps gcc libc-dev make\n\nWORKDIR /var/www/" + directory + "\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "CMD [\"gunicorn\", \"-b\", \"0.0.0.0:80\", \"main:app\"]"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)

					if err3 != nil {
						log.Fatal(err3)
					}
				default:
					os.Exit(0)
				}
				fmt.Print(e)
			default:
				continue
			}
		}
	case "typescript":
		js := [14]string{"react", "express", "angular", "vue", "ember", "@adonisjs/framework", "@hapi/hapi", "koa", "meteor", "feather", "sails", "nest", "socket", "derby"}
		for _, e := range js {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				path := directory
				switch e {
				case "adonis":
					fmt.Print(e)
				case "@hapi/hapi":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM node:lts-alpine\n\nRUN rm -rf /var/cache/apk/*\n\n\nCOPY . /var/www/" + directory + "\n\nENV PORT=8080\n\nRUN npm install\n\nEXPOSE $PORT\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "CMD [\"node\", \"server.js\"]"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)

					if err3 != nil {
						log.Fatal(err3)
					}
				case "express":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM node:lts-alpine\n\nCOPY . /var/www/" + directory + "\n\nENV PORT=3000\n\nWORKDIR /var/www/" + directory + "\n\nRUN npm install\n\nEXPOSE $PORT\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "ENTRYPOINT [\"node\", \"express.js\"]"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)

					if err3 != nil {
						log.Fatal(err3)
					}
				case "feather", "koa":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM node:lts-alpine\n\nCOPY . /var/www/" + directory + "\n\nENV PORT=3000\n\nRUN npm install\n\nEXPOSE $PORT\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "CMD [\"npm\", \"run\",  \"start\"]"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)

					if err3 != nil {
						log.Fatal(err3)
					}
				case "react":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM node:latest\n\nCOPY . /var/www/" + directory + "\n\nWORKDIR /var/www/" + directory + "\n\nENV PATH /var/www/" + directory + "/node_modules/.bin:$PATH\n\nRUN npm install --silent\n\nRUN npm install react-scripts@3.4.1 -g --silent\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "CMD [\"npm\", \"start\"]"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)

					if err3 != nil {
						log.Fatal(err3)
					}
				case "angular":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM node:latest\n\nCOPY . /var/www/" + directory + "\n\nWORKDIR  /var/www/" + directory + "\n\nRUN npm install -g @angular/cli\n\nRUN npm install\n\nEXPOSE 4200\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "CMD [\"npm\", \"run\", \"start\"]"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)

					if err3 != nil {
						log.Fatal(err3)
					}
				case "vue":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM node:lts-alpine\n\nRUN npm install -g http-server\n\nCOPY . /var/www/" + directory + "\n\nWORKDIR  ./var/www/" + directory + "\n\nENV PORT=3000\n\nRUN npm install\n\nRUN npm run build\n\nEXPOSE $PORT\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "CMD [\"http-server\", \"dist\"]"
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
				continue
			}
		}
	case "javascript":
		js := [14]string{"react", "express", "angular", "vue", "ember", "@adonisjs/framework", "@hapi/hapi", "koa", "meteor", "feather", "sails", "nest", "socket", "derby"}
		for _, e := range js {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				path := directory
				switch e {
				case "adonis":
					fmt.Print(e)
				case "@hapi/hapi":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM node:lts-alpine\n\nRUN rm -rf /var/cache/apk/*\n\n\nCOPY . /var/www/" + directory + "\n\nENV PORT=8080\n\nRUN npm install\n\nEXPOSE $PORT\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "CMD [\"node\", \"server.js\"]"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)

					if err3 != nil {
						log.Fatal(err3)
					}
				case "express":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM node:lts-alpine\n\nCOPY . /var/www/" + directory + "\n\nENV PORT=3000\n\nWORKDIR /var/www/" + directory + "\n\nRUN npm install\n\nEXPOSE $PORT\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "ENTRYPOINT [\"node\", \"express.js\"]"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)

					if err3 != nil {
						log.Fatal(err3)
					}
				case "feather", "koa":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM node:lts-alpine\n\nCOPY . /var/www/" + directory + "\n\nENV PORT=3000\n\nRUN npm install\n\nEXPOSE $PORT\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "CMD [\"npm\", \"run\",  \"start\"]"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)

					if err3 != nil {
						log.Fatal(err3)
					}
				case "react":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM node:latest\n\nCOPY . /var/www/" + directory + "\n\nWORKDIR /var/www/" + directory + "\n\nENV PATH /var/www/" + directory + "/node_modules/.bin:$PATH\n\nRUN npm install --silent\n\nRUN npm install react-scripts@3.4.1 -g --silent\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "CMD [\"npm\", \"start\"]"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)

					if err3 != nil {
						log.Fatal(err3)
					}
				case "angular":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM node:latest\n\nCOPY . /var/www/" + directory + "\n\nWORKDIR  /var/www/" + directory + "\n\nRUN npm install -g @angular/cli\n\nRUN npm install\n\nEXPOSE 4200\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "CMD [\"npm\", \"run\", \"start\"]"
					data2 := []byte(line2)

					var idx int64 = int64(len(data))

					_, err3 := f.WriteAt(data2, idx)

					if err3 != nil {
						log.Fatal(err3)
					}
				case "vue":
					if _, err := os.Stat(path); os.IsNotExist(err) {
						os.Mkdir(path, 0755)
					}
					f, err := os.Create(path + "/Dockerfile")

					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					firstLine := "FROM node:lts-alpine\n\nRUN npm install -g http-server\n\nCOPY . /var/www/" + directory + "\n\nWORKDIR  ./var/www/" + directory + "\n\nENV PORT=3000\n\nRUN npm install\n\nRUN npm run build\n\nEXPOSE $PORT\n\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)

					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "CMD [\"http-server\", \"dist\"]"
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
				continue
			}
		}
	case "php":
		php := [8]string{"laravel/framework", "symfony/symfony", "slim/slim", "zend", "phalcon", "cakephp/cakephp", "yii", "codeigniter"}
		for _, e := range php {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				fmt.Print(e)
			default:
				continue
			}
		}
	case "java":
		jv := [8]string{"spring", "play", "gwt", "hibernate", "struts", "vaadin", "yii", "codeigniter"}
		for _, e := range jv {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				fmt.Print(e)
			default:
				continue
			}
		}
	case "go":
		golang := [8]string{"gorilla", "echo", "gin", "beego", "fasthttp", "kit", "fiber", "iris"}
		for _, e := range golang {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				fmt.Print(e)
			default:
				continue
			}
		}
	case "elixir":
		golang := [9]string{"phoenix", "nerves", "sugar", "hedwig", "plug", "trot", "placid", "maru", "kitto"}
		for _, e := range golang {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				fmt.Print(e)
			default:
				continue
			}
		}
	case "ruby":
		rb := [10]string{"rubyonrails", "sinatra-activerecord", "roda", "camping", "cuba", "nancy", "remaze", "goliath", "hanami", "padrinho"}
		for _, e := range rb {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				fmt.Print(e)
			default:
				continue
			}
		}
	case "kotlin":
		rb := [8]string{"ktor", "kweb", "javalin", "spark", "spring", "jooby", "hexagon", "tekniq"}
		for _, e := range rb {
			switch contains(projeto.WhatsIsFrameworks(languagesProgram, directory), e) {
			case true:
				fmt.Print(e)
			default:
				continue
			}
		}
	case "c":
		fmt.Print("C")
	case "cplusplus":
		fmt.Print("C++")
	case "c#":
		fmt.Print("C#")
	case "rust":
		fmt.Print("Rust")
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

}
