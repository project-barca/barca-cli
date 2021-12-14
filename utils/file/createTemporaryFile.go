package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

func NewTempFile(tempFileName string, envVariable string) {
	tmpfile, _ := ioutil.TempFile("", tempFileName)
	fmt.Println(tmpfile.Name())
	os.Setenv(envVariable, tmpfile.Name())
}
