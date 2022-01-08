package dir

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func DiskUsage(path string, info os.FileInfo) int64 {
	size := info.Size()

	if !info.IsDir() {
		return size
	}

	dir, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return size
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.Name() == "." || file.Name() == ".." {
			continue
		}
		size += DiskUsage(path+"/"+file.Name(), file)
	}
	fmt.Printf("Tamanho em bytes: [%d] : [%s]\n", size, path)

	return size
}

func Scan(path string, f os.FileInfo, err error) error {
	//fmt.Printf("\n\nArquivo Verificado:     %s\nTamanho do arquivo:     %d bytes\n", path, f.Size())
	// string to write to file
	bytes := []byte(path + "\n")
	ENV_BARCA := os.Getenv("BARCA_TMP_FILE")
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
