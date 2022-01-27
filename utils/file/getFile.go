package file

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/crypto/md4"
)

var (
	fileInfo os.FileInfo
	err      error
	files    = make(map[[sha512.Size]byte]string)
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

	thepath, err := filepath.Abs(filedirectory)
	if err != nil {
		log.Fatal(err)
	}

	return thepath
}

func GetPermission(fileName string) {
	info, _ := os.Stat(fileName)
	mode := info.Mode()

	fmt.Println(fileName, "permissão: ", mode)
}

func CheckDuplicate(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if info.IsDir() { // skip directory
		return nil
	}
	data, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	hash := sha512.Sum512(data) // get the file sha512 hash

	if v, ok := files[hash]; ok {
		fmt.Printf("%q é uma duplicada de %q\n", path, v)
	} else {
		files[hash] = path // store in map for comparison
	}

	return nil
}

func GetHashSHA(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
}

func GetHashMD5(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
}

func GetHashMD4(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md4.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
}

// Verifique o tipo de conteúdo do arquivo específicado
func DetectContentType(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filetype := http.DetectContentType(buff)

	return filetype
}
