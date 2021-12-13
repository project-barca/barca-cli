package dir

import (
	"fmt"
	"os"
	"path/filepath"
)

func Scan(path string, f os.FileInfo, err error) error {
	fmt.Printf("\n\nArquivo Verificado:     %s\nTamanho do arquivo:     %d bytes\n", path, f.Size())
	return nil
}

func Scanner(directory string) {
	dir := directory

	err := filepath.Walk(dir, Scan)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
