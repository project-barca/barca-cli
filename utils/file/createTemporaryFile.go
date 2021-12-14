package file

import (
	"io/ioutil"
	"os"
)

func NewTempFile(tempFileName string, envVariable string) {
	tmpfile, _ := ioutil.TempFile("", tempFileName)
	os.Setenv(envVariable, tmpfile.Name())
}
