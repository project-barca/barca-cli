package dir

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

// dir.DiskUsage("/meu/diretorio", ".")
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
	// fmt.Printf("Tamanho em bytes: [%d] : [%s]\n", size, path)
	var sizeMB, sizeGB, sizeTB, sizePB, sizeEB = CalcBytes(size)

	if sizeMB != 0 {
		return sizeMB
	} else if sizeGB != 0 {
		return sizeGB
	} else if sizeTB != 0 {
		return sizeTB
	} else if sizePB != 0 {
		return sizePB
	} else if sizeEB != 0 {
		return sizeEB
	} else {
		return size
	}
}

func Scan(path string, f os.FileInfo, err error) error {
	//fmt.Printf("Coletando %s   |       tamnho:       %d bytes\n", path, f.Size())
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

// Find files by extension
//
// Example: dir.FilesByExtension("C:\\Windows", "exe")
func FilesByExt(directory, ext string) {
	dirname := directory + string(filepath.Separator)

	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == "."+ext {
				fmt.Println(file.Name())
			}
		}
	}
}

// Convert bytes in MB, GB, TB, PB, EB etc...
func CalcBytes(bytes int64) (mb, gb, tb, pb, eb int64) {
	// MegaByte
	sizeMB := bytes / (1 << (10 * 2))
	// GigaByte
	sizeGB := bytes / (1 << (10 * 3))
	// TeraByte
	sizeTB := bytes / (1 << (10 * 4))
	// PetaByte
	sizePB := bytes / (1 << (10 * 5))
	// ExaByte
	sizeEB := bytes / (1 << (10 * 6))
	// ZettaByte, YottaByte....

	return sizeMB, sizeGB, sizeTB, sizePB, sizeEB
}
