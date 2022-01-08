package file

import (
	"compress/gzip"
	"io/ioutil"
	"log"
	"os"
)

func Compress(inFile string) {
	// Open file for reading
	file, err := os.Open(inFile)
	if err != nil {
		log.Fatal(err)
	}
	// Read file content
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	// Close file
	file.Close()
	// Create .gz file to write to
	outputFile, err := os.Create(inFile + ".gz")
	if err != nil {
		log.Fatal(err)
	}
	// Create a gzip writer on top of file writer
	gzipWriter := gzip.NewWriter(outputFile)
	// Write out compressed data
	_, err = gzipWriter.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	// Close writer
	gzipWriter.Close()

	log.Println("Arquivo " + inFile + " compactado com sucesso!")
}
