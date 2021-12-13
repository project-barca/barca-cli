package dir

import (
	"io"
	"os"
)

func IsDirEmpty(directory string) bool {
	f, err := os.Open(directory)
	if err != nil {
		return false
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)

	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true
	}
	return false
}
