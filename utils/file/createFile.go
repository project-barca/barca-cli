package file

import (
	"fmt"
	"os"
)

func New(fileName string, destination string) {
	newFile, err := os.Create(destination + "/" + fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}
	newFile.Close()
}
