package integrate

import (
	"encoding/json"
	"log"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	scan "github.com/project-barca/barca-cli/scan/environment"
	"github.com/pterm/pterm"

	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func MySQL(directory string, dbservice string, dbname string, port string, host string, user string, password string, lang string) {
	// Configurações para internacionalização do software
	var isNode = scan.EnvironmentNode(directory)
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

	localizeConfigErrorMySQLExpress := i18n.LocalizeConfig{
		MessageID: "error_mysql_express",
	}
	localizeConfigSuccessMySQLExpress := i18n.LocalizeConfig{
		MessageID: "success_mysql_express",
	}
	resultErrorMySQLExpress, _ := localizer.Localize(&localizeConfigErrorMySQLExpress)
	resultSuccessMySQLExpress, _ := localizer.Localize(&localizeConfigSuccessMySQLExpress)

	if isNode == true {
		if os.IsNotExist(errorPath) {
			pterm.Error.Println(resultErrorMySQLExpress)
		} else {
			path := directory + "/config"

			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Mkdir(path, 0755)
			}
			f, err := os.Create(path + "/" + dbservice + ".js")

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			firstLine := "module.exports = {\n  HOST:  '" + host + "',\n  USER:  '" + user + "',\n  PASSWORD:  '" + password + "',\n  DB:  '" + dbname + "'\n"
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
	}
}
