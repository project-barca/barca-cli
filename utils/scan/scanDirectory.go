package scan

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"barca-cli/utils/dir"
	"barca-cli/utils/file"
)

func Directory(directory string) []string {
	file.NewTempFile("barco-cli.tmp.txt", "BARCA_TMP_FILE")
	dir.Scanner(directory)

	input, err := ioutil.ReadFile(os.Getenv("BARCA_TMP_FILE"))
	if err != nil {
		log.Println(err)
	}
	lines := strings.Split(string(input), "\n")

	return lines
}

func Extensions(directory string) []string {
	return file.WhatsIsExtension(Directory(directory))
}
