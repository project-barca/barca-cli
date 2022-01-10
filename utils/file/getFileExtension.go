package file

import (
	"path/filepath"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func WhatsIsExtension(fileName []string) []string {
	var extensions []string

	for i := 0; i < len(fileName); i++ {
		if Contains(extensions, filepath.Ext(fileName[i])) != true {
			extensions = append(extensions, filepath.Ext(fileName[i]))
		}
	}
	return extensions
}
