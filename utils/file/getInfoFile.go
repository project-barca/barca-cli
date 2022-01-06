package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	fileInfo os.FileInfo
	err      error
)

func GetInfo(fileName string) {

	fileInfo, err = os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nNome do Arquivo:", fileInfo.Name())
	fmt.Println("Tamanho:", fileInfo.Size())
	fmt.Println("Permissões:", fileInfo.Mode())
	fmt.Println("Últimas Modificações:", fileInfo.ModTime())
	fmt.Println("É Diretório: ", fileInfo.IsDir())

}

func CurrentPath(fileName string) string {
	filedirectory := filepath.Dir(fileName)

	path, err := filepath.Abs(filedirectory)
	if err != nil {
		log.Fatal(err)
	}

	return path
}
