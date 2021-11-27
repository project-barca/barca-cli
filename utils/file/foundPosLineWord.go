package file

import (
	"bufio"
	"os"
	"strings"

	"github.com/pterm/pterm"
)

func FindPositionLineText(path string, text string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	found := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), text) {
			pterm.Success.Println(found)
		}
		found++
	}

	if err := scanner.Err(); err != nil {
		return
	}

}
