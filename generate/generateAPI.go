package generate

import (
	"bufio"
	"encoding/json"
	"io"
	"time"

	"log"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	write "github.com/project-barca/barca-cli/generate/node/write/js"
	"github.com/pterm/pterm"
	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func API(lang string, directory string, framework string, port string, projectName string) {
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

	switch framework {
	case "express":
		if directory == "" {

		}
		write.ServerExpress(port, directory, lang)
	case "echo":

	case "gorilla":
		if directory == "" {
		}
		server, err := os.Open("frameworks/http/gorilla/gorilla-mux.go")
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := server.Close(); err != nil {
				panic(err)
			}
		}()
		r := bufio.NewReader(server)

		_, errorPath := os.Stat("./" + directory)

		localizeConfigWelcome := i18n.LocalizeConfig{
			MessageID: "generate_api",
		}
		localizeConfigCreateDirectory := i18n.LocalizeConfig{
			MessageID: "create_directory_api",
		}
		localizeConfigErrorDirectory := i18n.LocalizeConfig{
			MessageID: "error_directory_api",
		}
		localizeConfigSuccessProject := i18n.LocalizeConfig{
			MessageID: "create_api_success",
		}
		resultWelcome, _ := localizer.Localize(&localizeConfigWelcome)
		resultDir, _ := localizer.Localize(&localizeConfigCreateDirectory)
		resultErrorDir, _ := localizer.Localize(&localizeConfigErrorDirectory)
		resultProjectSuccess, _ := localizer.Localize(&localizeConfigSuccessProject)

		if os.IsNotExist(errorPath) {

			spinnerSuccess, _ := pterm.DefaultSpinner.Start(resultWelcome)
			time.Sleep(time.Second * 2)
			spinnerSuccess.Success()

			errDir := os.Mkdir("./"+directory, 0777)
			if errDir != nil {
				log.Fatal(errorPath)
				return
			}
			currentServer, err := os.Create(projectName + "./main.go")

			spinnerSuccessDir, _ := pterm.DefaultSpinner.Start(resultDir)
			time.Sleep(time.Second * 2)
			spinnerSuccessDir.Success(resultProjectSuccess)
			time.Sleep(time.Second * 1)
			spinnerSuccessDir.Success("OK")

			if err != nil {
				panic(err)
			}
			defer func() {
				if err := currentServer.Close(); err != nil {
					panic(err)
				}
			}()

			w := bufio.NewWriter(currentServer)

			buf := make([]byte, 1024)
			for {
				n, err := r.Read(buf)
				if err != nil && err != io.EOF {
					panic(err)
				}
				if n == 0 {
					break
				}
				if _, err := w.Write(buf[:n]); err != nil {
					panic(err)
				}
			}

			if err = w.Flush(); err != nil {
				panic(err)
			}
		} else {
			pterm.Error.Println(resultErrorDir)

		}

	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

}
