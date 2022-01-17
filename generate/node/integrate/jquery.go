package integrate

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pterm/pterm"
	"golang.org/x/text/language"
)

func JQuery(directory, lang string) {

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

	localizeSuccessJQuery := i18n.LocalizeConfig{
		MessageID: "success_jquery",
	}
	localizeDownloadModulesJQuery := i18n.LocalizeConfig{
		MessageID: "download_modules_jquery",
	}

	resultJQueryServer, _ := localizer.Localize(&localizeSuccessJQuery)
	resultDownloadModulesJQuery, _ := localizer.Localize(&localizeDownloadModulesJQuery)

	_, errorPath := os.Stat("./" + directory)

	if os.IsNotExist(errorPath) {
		errDir := os.Mkdir("./"+directory, 0777)
		if errDir != nil {
			log.Fatal(errorPath)
			return
		}

		f, err := os.Create("./" + directory + "/main.js")
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		line := "$(document).ready(function() {\n\n});"
		data := []byte(line)

		_, err2 := f.Write(data)
		if err2 != nil {
			log.Fatal(err2)
		}

		spinnerSuccess, _ := pterm.DefaultSpinner.Start(resultDownloadModulesJQuery)
		time.Sleep(time.Second * 5)
		spinnerSuccess.Success()

		pterm.Success.Println(resultJQueryServer)
	} else {
		f, err := os.Create("./" + directory + "/main.js")
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		line := "$(document).ready(function() {\n\n});"
		data := []byte(line)

		_, err2 := f.Write(data)
		if err2 != nil {
			log.Fatal(err2)
		}

		spinnerSuccess, _ := pterm.DefaultSpinner.Start(resultDownloadModulesJQuery)
		time.Sleep(time.Second * 5)
		spinnerSuccess.Success()

		pterm.Success.Println(resultJQueryServer)
	}
}
