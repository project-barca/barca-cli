package file

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func File2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// Insert sting to n-th line of file.
// If you want to insert a line, append newline '\n' to the end of the string.
func InsertStringToFile(path, str string, index int) error {
	lines, err := File2lines(path)
	if err != nil {
		return err
	}

	fileContent := ""
	for i, line := range lines {
		if i == index {
			fileContent += str
		}
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}

func GetStringLinesByFile(path string) []string {
	lines, err := File2lines(path)
	if err != nil {
		return nil
	}
	return lines
}

func GetStringLineByIndex(path string, index int) string {
	lines, err := File2lines(path)
	var output string
	if err != nil {
		fmt.Errorf("Error")
	}

	for i, line := range lines {
		if i == index {
			output = line
		}
	}
	return output
}

func GetNumberLines(path string) int {
	lines, err := File2lines(path)
	if err != nil {
		fmt.Errorf("Error")
	}
	return len(lines)
}

func ReplaceText(inFile string, oldWord string, newWord string) {
	input, err := ioutil.ReadFile(inFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte(oldWord), []byte(newWord), -1)

	if err = ioutil.WriteFile("new-"+inFile, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ConvertToString(inFile string) string {
	file, err := ioutil.ReadFile(inFile)
	if err != nil {
		fmt.Println(err)
	}
	// out the file content
	fmt.Println(string(file))
	return string(file)
}
