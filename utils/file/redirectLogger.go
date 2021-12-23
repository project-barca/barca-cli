package file

import (
	"log"
	"os"
)

func RedirectLogFile(content string, nameFile string) {
	file, e := os.OpenFile(nameFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatalln("Falha ao tentar abrir o arquivo")
	}

	log.SetOutput(file)
	log.Println(content)
}
