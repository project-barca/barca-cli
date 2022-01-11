package dir

import (
	"fmt"
	"io"
	"os"
)

func IsDirEmpty(directory string) bool {
	f, err := os.Open(directory)
	if err != nil {
		return false
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)

	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true
	}
	return false
}

func IsDir(path string) bool {
	object := "./" + path

	fdir, err := os.Open(object)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fdir.Close()

	finfo, err := fdir.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var res bool
	switch mode := finfo.Mode(); {
	case mode.IsDir():
		fmt.Println("Objeto é um diretório")
		res = true
	case mode.IsRegular():
		fmt.Println("Objeto é um arquivo")
		res = false
	}
	return res
}
