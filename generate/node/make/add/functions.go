package add

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/project-barca/barca-cli/utils/file"
	"github.com/pterm/pterm"

	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func Function(lang, directory, collection, database, method string, hidden bool, inputs []string) {
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

	_, errorPath := os.Stat("./" + directory)
	switch method {
	case "create":
		localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
			MessageID: "error_functions_javascript",
		}
		localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
			MessageID: "success_functions_javascript",
		}
		resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
		resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

		if os.IsNotExist(errorPath) {
			pterm.Error.Println(resultErrorModelFunctionsJavascript)
		} else {
			path := directory + "/js"

			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Mkdir(path, 0755)
			}

			fmt.Print(file.CheckIfExists(path + "/" + collection + ".js"))
			if file.CheckIfExists(path+"/"+collection+".js") != false {
				indexLineFile := file.FindPositionLineText(path+"/"+collection+".js", "function insert"+collection+"() {")
				fmt.Println(indexLineFile)
				fmt.Println(file.GetNumberLines(path + "/" + collection + ".js"))
				if indexLineFile == 0 {
					functiondata := []string{"\nfunction insert" + collection + "() {",
						"};",
					}
					file, err := os.OpenFile(path+"/"+collection+".js", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
					}

					datawriter := bufio.NewWriter(file)

					for _, data := range functiondata {
						_, _ = datawriter.WriteString(data + "\n")
					}

					datawriter.Flush()
					file.Close()
				} else {
					fmt.Println("Já existe função para criar")
				}
			} else {
				f, err := os.Create(path + "/" + collection + ".js")
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()

				firstLine := "function insert" + collection + "() {"
				data := []byte(firstLine)

				_, err2 := f.Write(data)

				if err2 != nil {
					log.Fatal(err2)
				}

				line2 := "\n" + inputs[2] + "\n};"
				data2 := []byte(line2)

				var idx int64 = int64(len(data))

				_, err3 := f.WriteAt(data2, idx)
				if err3 != nil {
					log.Fatal(err3)
				}
			}

			pterm.Success.Println(resultSuccessFunctionsJavascript)
		}
	case "read":
		localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
			MessageID: "error_functions_javascript",
		}
		localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
			MessageID: "success_functions_javascript",
		}
		resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
		resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

		if os.IsNotExist(errorPath) {
			pterm.Error.Println(resultErrorModelFunctionsJavascript)
		} else {
			path := directory + "/js"

			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Mkdir(path, 0755)
			}

			fmt.Print(file.CheckIfExists(path + "/" + collection + ".js"))
			if file.CheckIfExists(path+"/"+collection+".js") != false {
				indexLineFile := file.FindPositionLineText(path+"/"+collection+".js", "function read"+collection+"() {")
				fmt.Println(indexLineFile)
				fmt.Println(file.GetNumberLines(path + "/" + collection + ".js"))
				if indexLineFile == 0 {
					functiondata := []string{"\nfunction read" + collection + "() {",
						"};",
					}
					file, err := os.OpenFile(path+"/"+collection+".js", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
					}

					datawriter := bufio.NewWriter(file)

					for _, data := range functiondata {
						_, _ = datawriter.WriteString(data + "\n")
					}

					datawriter.Flush()
					file.Close()
				} else {
					fmt.Println("Já existe função para ler")
				}
			} else {
				f, err := os.Create(path + "/" + collection + ".js")
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()

				firstLine := "function read" + collection + "() {"
				data := []byte(firstLine)

				_, err2 := f.Write(data)

				if err2 != nil {
					log.Fatal(err2)
				}

				line2 := "\n" + inputs[2] + "\n};"
				data2 := []byte(line2)

				var idx int64 = int64(len(data))

				_, err3 := f.WriteAt(data2, idx)
				if err3 != nil {
					log.Fatal(err3)
				}
			}

			pterm.Success.Println(resultSuccessFunctionsJavascript)
		}
	case "delete":
		localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
			MessageID: "error_functions_javascript",
		}
		localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
			MessageID: "success_functions_javascript",
		}
		resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
		resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

		if os.IsNotExist(errorPath) {
			pterm.Error.Println(resultErrorModelFunctionsJavascript)
		} else {
			path := directory + "/js"

			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Mkdir(path, 0755)
			}

			fmt.Print(file.CheckIfExists(path + "/" + collection + ".js"))
			if file.CheckIfExists(path+"/"+collection+".js") != false {
				indexLineFile := file.FindPositionLineText(path+"/"+collection+".js", "function delete"+collection+"() {")
				fmt.Println(indexLineFile)
				fmt.Println(file.GetNumberLines(path + "/" + collection + ".js"))
				if indexLineFile == 0 {
					functiondata := []string{"\nfunction delete" + collection + "() {",
						"};",
					}
					file, err := os.OpenFile(path+"/"+collection+".js", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
					}

					datawriter := bufio.NewWriter(file)

					for _, data := range functiondata {
						_, _ = datawriter.WriteString(data + "\n")
					}

					datawriter.Flush()
					file.Close()
				} else {
					fmt.Println("Já existe função para excluír")
				}
			} else {
				f, err := os.Create(path + "/" + collection + ".js")
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()

				firstLine := "function delete" + collection + "() {"
				data := []byte(firstLine)

				_, err2 := f.Write(data)

				if err2 != nil {
					log.Fatal(err2)
				}

				line2 := "\n" + inputs[2] + "\n};"
				data2 := []byte(line2)

				var idx int64 = int64(len(data))

				_, err3 := f.WriteAt(data2, idx)
				if err3 != nil {
					log.Fatal(err3)
				}
			}

			pterm.Success.Println(resultSuccessFunctionsJavascript)
		}
	case "update":
		localizeConfigErrorModelFunctionsJavacript := i18n.LocalizeConfig{
			MessageID: "error_functions_javascript",
		}
		localizeConfigSuccessFunctionsJavascript := i18n.LocalizeConfig{
			MessageID: "success_functions_javascript",
		}
		resultErrorModelFunctionsJavascript, _ := localizer.Localize(&localizeConfigErrorModelFunctionsJavacript)
		resultSuccessFunctionsJavascript, _ := localizer.Localize(&localizeConfigSuccessFunctionsJavascript)

		if os.IsNotExist(errorPath) {
			pterm.Error.Println(resultErrorModelFunctionsJavascript)
		} else {
			path := directory + "/js"

			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Mkdir(path, 0755)
			}

			fmt.Print(file.CheckIfExists(path + "/" + collection + ".js"))
			if file.CheckIfExists(path+"/"+collection+".js") != false {
				indexLineFile := file.FindPositionLineText(path+"/"+collection+".js", "function update"+collection+"() {")
				fmt.Println(indexLineFile)
				fmt.Println(file.GetNumberLines(path + "/" + collection + ".js"))
				if indexLineFile == 0 {
					functiondata := []string{"\nfunction update" + collection + "() {",
						"};",
					}
					file, err := os.OpenFile(path+"/"+collection+".js", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
					}

					datawriter := bufio.NewWriter(file)

					for _, data := range functiondata {
						_, _ = datawriter.WriteString(data + "\n")
					}

					datawriter.Flush()
					file.Close()
				} else {
					fmt.Println("Já existe função para atualizar")
				}
			} else {
				f, err := os.Create(path + "/" + collection + ".js")
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()

				firstLine := "function update" + collection + "() {"
				data := []byte(firstLine)

				_, err2 := f.Write(data)

				if err2 != nil {
					log.Fatal(err2)
				}

				line2 := "\n" + inputs[2] + "\n};"
				data2 := []byte(line2)

				var idx int64 = int64(len(data))

				_, err3 := f.WriteAt(data2, idx)
				if err3 != nil {
					log.Fatal(err3)
				}
			}

			pterm.Success.Println(resultSuccessFunctionsJavascript)
		}
	default:
		// barcaSQL
	}
}
