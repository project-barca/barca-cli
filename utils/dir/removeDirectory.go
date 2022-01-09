package dir

import (
	"os"
)

func RemoveDir(path string) {
	os.Remove(path)
}

func RemoveAllDir(path string) {
	os.RemoveAll(path)
}
