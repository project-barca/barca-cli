package file

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
)

func RemoveFile(fileName string) {
	var targetFile = fileName

	file, err := os.OpenFile(targetFile, os.O_RDWR, 0666)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	// find out how large is the target file
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	// calculate the new slice size
	// base on how large our target file is
	var size int64 = fileInfo.Size()
	zeroBytes := make([]byte, size)

	// fill out the new slice with 0 value
	copy(zeroBytes[:], "0")
	fmt.Println(zeroBytes[:])

	// wipe the content of the target file
	n, err := file.Write([]byte(zeroBytes))

	if err != nil {
		panic(err)
	}
	fmt.Printf("Wiped %v bytes.\n", n)

	// finally remove/delete our file
	err = os.Remove(targetFile)

	if err != nil {
		panic(err)
	}
}

func RemoveBigFile(fileName string) {
	var targetFile = fileName

	file, err := os.OpenFile(targetFile, os.O_RDWR, 0666)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	// find out how large is the target file
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	// calculate the new slice size
	// base on how large our target file is
	var fileSize int64 = fileInfo.Size()
	const fileChunk = 1 * (1 << 20) // 1 MB, change this to your requirement

	// calculate total number of parts the file will be chunked into
	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	lastPosition := 0

	for i := uint64(0); i < totalPartsNum; i++ {
		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partZeroBytes := make([]byte, partSize)
		// fill out the part with zero value
		copy(partZeroBytes[:], "0")
		// over write every byte in the chunk with 0
		fmt.Println("Writing to target file from position : ", lastPosition)
		n, err := file.WriteAt([]byte(partZeroBytes), int64(lastPosition))
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Wiped %v bytes.\n", n)
		// update last written position
		lastPosition = lastPosition + partSize
	}
	// finally remove/delete our file
	err = os.Remove(targetFile)
	if err != nil {
		panic(err)
	}
}

func RemoveFilesByExtension(path string, ext string) {
	dirname := "." + string(filepath.Separator) + path

	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ext {
				os.Remove(path + string(filepath.Separator) + file.Name())
				fmt.Println("Deletado ", file.Name())
			}
		}
	}
}
