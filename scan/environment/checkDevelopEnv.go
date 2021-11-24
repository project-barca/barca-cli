package scan

import (
	"fmt"
	"os"
)

func EnvironmentNode(directory string) bool {
	// get FileInfo structure describing file
	fileInfo, err := os.Stat(directory + "/package.json")
	// check if there is an error
	isNode := false
	if err != nil {
		// check if error is file does not exist
		if os.IsNotExist(err) {
			isNode = false
			fmt.Println("Não existe nenhuma possibilidade para identificarmos que tipo de projeto está trabalhando\n\nExemplos: PHP, Node.js, Python, Go.")
			return isNode
		}
	} else {
		fmt.Println(fileInfo.Name(), "existe")
		isNode = true
		return isNode
	}
	return isNode
}
