package file

import (
	"fmt"
	"io/ioutil"
)

func NewTempFile(tempFileName string) {
	tmpfile, _ := ioutil.TempFile("", tempFileName)
	fmt.Println(tmpfile.Name())
}
