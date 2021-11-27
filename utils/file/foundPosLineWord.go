package file

import (
	"bufio"
	"os"
	"strings"
)

func FindPositionLineText(path string, text string) int {
	var foundPos int
	f, err := os.Open(path)
	if err != nil {

	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	found := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), text) {
			foundPos = found
		}
		found++
	}

	if err := scanner.Err(); err != nil {
		return 0
	}
	return foundPos
}
