package dependencies

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/project-barca/barca-cli/utils/file"
	"github.com/pterm/pterm"

	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func PackageJson(project, main, lang string) {
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

	_, errorPath := os.Stat("./" + project)

	localizeConfigErrorPackageJson := i18n.LocalizeConfig{
		MessageID: "error_package_json",
	}
	localizeConfigSuccessPackageJson := i18n.LocalizeConfig{
		MessageID: "success_package_json",
	}
	resultErrorPackageJson, _ := localizer.Localize(&localizeConfigErrorPackageJson)
	resultSuccessPackageJson, _ := localizer.Localize(&localizeConfigSuccessPackageJson)

	if os.IsNotExist(errorPath) {
		pterm.Error.Println(resultErrorPackageJson)
	} else {

		f, err := os.Create("./" + project + "/package.json")

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		firstLine := "{"
		data := []byte(firstLine)

		_, err2 := f.Write(data)

		if err2 != nil {
			log.Fatal(err2)
		}

		line2 := "\n  \"name\":  \"" + project + "\",\n  \"version\":	\"1.0.0\",\n	\"main\": \"express.js\",\n  \"scripts\": {\n    \"start\": \"node express.js\"\n  },\n  \"dependencies\":  {\n    \"express\": \"~4.17.1\"\n	}\n}"
		data2 := []byte(line2)

		var idx int64 = int64(len(data))

		_, err3 := f.WriteAt(data2, idx)

		if err3 != nil {
			log.Fatal(err3)
		}

		pterm.Success.Println(resultSuccessPackageJson)
	}
}

func PackagePip(project, main, lang, framework string) {
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

	_, errorPath := os.Stat("./" + project)

	localizeConfigErrorPackageJson := i18n.LocalizeConfig{
		MessageID: "error_package_json",
	}
	localizeConfigSuccessPackageJson := i18n.LocalizeConfig{
		MessageID: "success_package_json",
	}
	resultErrorPackageJson, _ := localizer.Localize(&localizeConfigErrorPackageJson)
	resultSuccessPackageJson, _ := localizer.Localize(&localizeConfigSuccessPackageJson)

	if os.IsNotExist(errorPath) {
		pterm.Error.Println(resultErrorPackageJson)
	} else {

		f, err := os.Create("./" + project + "/requirements.txt")
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		switch framework {
		case "bottle":
			firstLine := "bottle==0.13.0\n"
			data := []byte(firstLine)

			_, err2 := f.Write(data)

			if err2 != nil {
				log.Fatal(err2)
			}

			line2 := "requests==2.27.1\n"
			data2 := []byte(line2)

			var idx int64 = int64(len(data))

			_, err3 := f.WriteAt(data2, idx)

			if err3 != nil {
				log.Fatal(err3)
			}
		case "flask":
			firstLine := "Flask==2.0.2\n"
			data := []byte(firstLine)

			_, err2 := f.Write(data)

			if err2 != nil {
				log.Fatal(err2)
			}

			line2 := "requests==2.27.1\n"
			data2 := []byte(line2)

			var idx int64 = int64(len(data))

			_, err3 := f.WriteAt(data2, idx)

			if err3 != nil {
				log.Fatal(err3)
			}
		default:
			// freebsd, openbsd,
		}

		pterm.Success.Println(resultSuccessPackageJson)
	}
}

// Get version of some module npm
func GetVersionModule(path string, module string) string {
	var version []string
	for i := 0; i < file.GetNumberLines(path+"/package.json"); i++ {
		if strings.Contains(file.GetStringLineByIndex(path+"/package.json", i), "\""+module+"\"") != false {
			version = strings.Split(file.GetStringLineByIndex(path+"/package.json", i), ":")
		}
	}
	return version[1]
}

// Get version of some module composer
func GetVersionModuleComposer(path string, module string) string {
	var version []string
	for i := 0; i < file.GetNumberLines(path+"/composer.json"); i++ {
		if strings.Contains(file.GetStringLineByIndex(path+"/composer.json", i), "\""+module+"\"") != false {
			version = strings.Split(file.GetStringLineByIndex(path+"/composer.json", i), ":")
		}
	}
	return version[1]
}

// Check modules
func IfExistsModule(path, language, module string) bool {
	var res bool

	switch language {
	case "javascript":
		for i := 0; i < file.GetNumberLines(path+"/package.json"); i++ {
			if strings.Contains(file.GetStringLineByIndex(path+"/package.json", i), "\""+module+"\"") != false {
				res = true
			}
		}
	case "python":
		for i := 0; i < file.GetNumberLines(path+"/requirements.txt"); i++ {
			if strings.Contains(file.GetStringLineByIndex(path+"/requirements.txt", i), module) != false {
				res = true
			}
		}
	case "php":
		for i := 0; i < file.GetNumberLines(path+"/composer.son"); i++ {
			if strings.Contains(file.GetStringLineByIndex(path+"/composer.json", i), "\""+module+"\"") != false {
				res = true
			}
		}
	default:
		os.Exit(0)
	}
	return res
}
