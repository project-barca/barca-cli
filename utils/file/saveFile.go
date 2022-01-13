package file

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func SaveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func SaveXML(fileName string, key interface{}) {
	filename := fileName
	file, err := os.Create(filename)
	checkError(err)
	xmlWriter := io.Writer(file)
	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	err = enc.Encode(key)
	checkError(err)
}

func SaveCSV(fileName string, key [][]string) {
	filename := fileName
	file, err := os.Create(filename)
	checkError(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range key {
		err := writer.Write(value)
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
