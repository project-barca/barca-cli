package file

import (
	"fmt"
	"os"
)

func CheckIfExists(fileName string) bool {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Arquivo não existe.")
			return false
		}
	}
	fmt.Println(fileInfo.Name(), "existe")
	return true
}
