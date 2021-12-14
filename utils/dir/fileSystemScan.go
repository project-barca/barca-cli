package dir

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func Scan(path string, f os.FileInfo, err error) error {
	//fmt.Printf("\n\nArquivo Verificado:     %s\nTamanho do arquivo:     %d bytes\n", path, f.Size())
	// string to write to file
	bytes := []byte(path + "\n")
	ENV_BARCA := os.Getenv("BARCA_TEMP_LIST_DIR")
	// open file for writing
	fd1, err := syscall.Open(ENV_BARCA, syscall.O_CREAT|syscall.O_RDWR|syscall.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Err:", err)
	}
	// write to file
	syscall.Write(fd1, bytes)
	// check results
	if err == nil {
		return nil
	} else {
		return err
	}

}

func Scanner(directory string) {
	dir := directory
	err := filepath.Walk(dir, Scan)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
