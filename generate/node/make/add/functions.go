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

func init() {
	log.SetPrefix("FUNCTIONS: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func FunctionJS(lang, directory, collection, database, method string, hidden bool, inputs []string) {
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

	if hidden != true {

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

				if file.CheckIfExists(path+"/"+collection+".js") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".js", "function insert"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfunction insert" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "  "+input+",")
							}
						}
						functiondata = append(functiondata, "};")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".js", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".js", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função CREATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".js")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "function insert" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("  " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".js", "};", file.GetNumberLines(path+"/"+collection+".js")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".js", "insert"+collection+"()", "insert"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
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

				if file.CheckIfExists(path+"/"+collection+".js") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".js", "function read"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfunction read" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "  "+input+",")
							}
						}
						functiondata = append(functiondata, "};")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".js", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".js", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função READ")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".js")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "function read" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("  " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".js", "};", file.GetNumberLines(path+"/"+collection+".js")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".js", "read"+collection+"()", "read"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
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

				if file.CheckIfExists(path+"/"+collection+".js") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".js", "function delete"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfunction delete" + collection + "(id) {",
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
						fmt.Println("Já existe está função DELETE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".js")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "function delete" + collection + "(id) {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}

					line2 := "id, \n};"
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

				if file.CheckIfExists(path+"/"+collection+".js") != false {
					indexLineFile := file.FindPositionLineText(path+"/"+collection+".js", "function update"+collection+"(")
					if indexLineFile == 0 {
						functiondata := []string{"\nfunction update" + collection + "() {"}
						var strParamns string
						if inputs != nil {
							for _, input := range inputs {
								strParamns += input + ", "
								functiondata = append(functiondata, "  "+input+",")
							}
						}
						functiondata = append(functiondata, "};")

						fileCurrent, err := os.OpenFile(path+"/"+collection+".js", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatalf("Falha ao tentar modificar o arquivo: %s", err)
						}

						datawriter := bufio.NewWriter(fileCurrent)
						for _, data := range functiondata {
							_, _ = datawriter.WriteString(data + "\n")
						}
						datawriter.Flush()
						fileCurrent.Close()

						file.ReplaceString(path+"/"+collection+".js", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					} else {
						fmt.Println("Já existe está função UPDATE")
					}
				} else {
					f, err := os.Create(path + "/" + collection + ".js")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					firstLine := "function update" + collection + "() {\n"
					data := []byte(firstLine)

					_, err2 := f.Write(data)
					if err2 != nil {
						log.Fatal(err2)
					}
					// Verify paramns input values - start
					if inputs != nil {
						var strParamns string
						for _, input := range inputs {
							strParamns += input + ", "
							_, err3 := f.WriteString("  " + input + ",\n")
							if err3 != nil {
								log.Fatal(err3)
							}
						}
						f.WriteString("\n")

						errStr := file.InsertStringToFile(path+"/"+collection+".js", "};", file.GetNumberLines(path+"/"+collection+".js")-1)
						if errStr != nil {
							log.Fatal(errStr)
						}

						file.ReplaceString(path+"/"+collection+".js", "update"+collection+"()", "update"+collection+"("+strParamns+")")
					}
					// Verify paramns input values - end
				}

				pterm.Success.Println(resultSuccessFunctionsJavascript)
			}
		default:
			// barcaSQL
		}
	}

}
