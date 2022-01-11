package file

import (
	"archive/zip"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Decompress(inFile string) {
	var ext = filepath.Ext(inFile)
	switch ext {
	case ".zip":
		zipfile := inFile
		if zipfile == "" {
			fmt.Println("Use: barca decompress sourcefile.zip")
			os.Exit(1)
		}

		reader, err := zip.OpenReader(zipfile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer reader.Close()

		for _, f := range reader.Reader.File {
			zipped, err := f.Open()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer zipped.Close()

			// get the individual file name and extract the current directory
			path := filepath.Join("./", f.Name)

			if f.FileInfo().IsDir() {
				os.MkdirAll(path, f.Mode())
				fmt.Println("Creating directory", path)
			} else {
				writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, f.Mode())
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				defer writer.Close()

				if _, err = io.Copy(writer, zipped); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				fmt.Println("Descompactado: ", path)
			}
		}

	case ".gzip":
		// Add gzip
		filename := inFile

		if filename == "" {
			fmt.Println("Utilize: barca decompress sourcefile.gzip")
			os.Exit(1)
		}
	case "zlib":
		filename := inFile

		if filename == "" {
			fmt.Println("Use: barca decompress sourcefile.zlib")
			os.Exit(1)
		}

		zlibfile, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		reader, err := zlib.NewReader(zlibfile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer reader.Close()

		newfilename := strings.TrimSuffix(filename, ".zlib")

		writer, err := os.Create(newfilename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer writer.Close()

		if _, err = io.Copy(writer, reader); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Descompactado: ", newfilename)
	}

}
