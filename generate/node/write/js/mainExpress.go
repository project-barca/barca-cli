package write

import (
	dependencies "barca-cli/generate/node/dependencies"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pterm/pterm"
	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func ServerExpress(port string, directory string, lang string) {
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

	localizeSuccessExpressServer := i18n.LocalizeConfig{
		MessageID: "success_server_express",
	}
	localizeSuccessNodeModules := i18n.LocalizeConfig{
		MessageID: "success_node_modules",
	}
	localizeInstallModulesNpm := i18n.LocalizeConfig{
		MessageID: "install_modules_npm",
	}
	resultInstallNpmModules, _ := localizer.Localize(&localizeInstallModulesNpm)
	resultExpressServer, _ := localizer.Localize(&localizeSuccessExpressServer)
	resultNodeModules, _ := localizer.Localize(&localizeSuccessNodeModules)

	_, errorPath := os.Stat("./" + directory)

	if os.IsNotExist(errorPath) {
		errDir := os.Mkdir("./"+directory, 0777)
		if errDir != nil {
			log.Fatal(errorPath)
			return
		}

		f, err := os.Create("./" + directory + "/express.js")

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		line := "const express = require('express'),"
		data := []byte(line)

		_, err2 := f.Write(data)

		if err2 != nil {
			log.Fatal(err2)
		}

		line2 := "\napp = express(" + port + ")\n"
		data2 := []byte(line2)

		var idx int64 = int64(len(data))

		_, err3 := f.WriteAt(data2, idx)

		if err3 != nil {
			log.Fatal(err3)
		}
		pterm.Success.Println(resultExpressServer)
		dependencies.PackageJson(directory, "express.js", lang)
		dependencies.AutoInstallNpm(directory)
		//go file.Move(directory, ".bat")
		curr_wd, err := os.Getwd()

		spinnerSuccess, _ := pterm.DefaultSpinner.Start(resultInstallNpmModules)
		time.Sleep(time.Second * 5)
		spinnerSuccess.Success()
		//cmd := exec.Command("cmd", "start", "/c", curr_wd+"\\install.bat")
		cmd := exec.Command("cmd", "start", "/c", curr_wd+"\\installMod.bat")
		if err := cmd.Start(); err != nil {
			fmt.Println("Error: ", err)
		}
		//fmt.Println(curr_wd + "\\" + directory)
		//curr_wd, err := os.Getwd()
		//cmd := exec.Command("cd", curr_wd+"/hao", "&", "npm", "run", "install")
		cmd.Wait()
		stdout, err := cmd.Output()
		pterm.Success.Println(resultNodeModules)
		//fmt.Println(curr_wd)
		fmt.Print(stdout)

	}
}
