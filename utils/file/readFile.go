package file

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadXML(fileName string, key interface{}) interface{} {
	data, _ := ioutil.ReadFile(fileName)

	interfaceKey := key
	_ = xml.Unmarshal([]byte(data), &interfaceKey)

	return interfaceKey
}

func ReadJSON(fileName string, key interface{}) interface{} {
	data, _ := ioutil.ReadFile(fileName)

	interfaceKey := key
	_ = json.Unmarshal([]byte(data), &interfaceKey)

	return interfaceKey
}

func ReadCSV(fileName string) []string {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()

	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}

	return txtlines
}
