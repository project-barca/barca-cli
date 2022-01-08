package file

import (
	"archive/zip"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func Compress(inFile string, ext string) {
	switch ext {
	case "gz":
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
	case "zip":
		newfile, err := os.Create(inFile + ".zip")
		if err != nil {
			log.Fatal(err)
		}
		defer newfile.Close()

		zipit := zip.NewWriter(newfile)

		defer zipit.Close()

		zipfile, err := os.Open(inFile)
		if err != nil {
			log.Fatal(err)
		}
		defer zipfile.Close()

		// get the file information
		info, err := zipfile.Stat()
		if err != nil {
			log.Fatal(err)
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			log.Fatal(err)
		}
		header.Method = zip.Deflate
		writer, err := zipit.CreateHeader(header)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(writer, zipfile)
		log.Println("Arquivo " + inFile + " compactado com sucesso!")
	}

}
