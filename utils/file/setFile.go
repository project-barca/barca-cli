package file

import (
	"os"
)

func CanWrite(filepath string) (bool, error) {
	file, err := os.OpenFile(filepath, os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			return false, err
		}
	}
	file.Close()
	return true, nil
}

func SetWritable(filepath string) error {
	err := os.Chmod(filepath, 0222)
	return err
}

func SetReadOnly(filepath string) error {
	err := os.Chmod(filepath, 0444)
	return err
}
