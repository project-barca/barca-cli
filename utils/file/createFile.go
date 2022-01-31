package file

import (
	"log"
	"os"
)

func New(fileName string, destination string) {
	newFile, err := os.Create(destination + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()
}

func NewScript(cmd, fileName, destination string) string {

	f, err := os.Create(destination + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	line := cmd
	data := []byte(line)

	_, err2 := f.Write(data)
	if err2 != nil {
		log.Fatal(err2)
	}

	return "Script criado: " + fileName
}
