package dependencies

import (
	"log"
	"os"
)

func AutoInstallNpm(project string) {
	_, errorPath := os.Stat("./" + project)

	if os.IsNotExist(errorPath) {

	} else {

		f, err := os.Create("./installMod.bat")

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		firstLine := "@ECHO OFF\n\ncd " + project
		data := []byte(firstLine)

		_, err2 := f.Write(data)

		if err2 != nil {
			log.Fatal(err2)
		}

		line2 := "\nnpm install"
		data2 := []byte(line2)

		var idx int64 = int64(len(data))

		_, err3 := f.WriteAt(data2, idx)

		if err3 != nil {
			log.Fatal(err3)
		}
	}
}
