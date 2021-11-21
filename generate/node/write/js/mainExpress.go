package write

import (
	dependencies "barca-cli/generate/node/dependencies"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func ServerExpress(port string, directory string, lang string) {
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
		go dependencies.PackageJson(directory, "express.js", lang)

		cmd := exec.Command("npm", "install", directory)
		curr_wd, err := os.Getwd()
		//cmd := exec.Command("cd", curr_wd+"/hao", "&", "npm", "run", "install")
		stdout, err := cmd.Output()
		fmt.Println("done")
		fmt.Println(curr_wd)
		fmt.Print(string(stdout))

	}
}
