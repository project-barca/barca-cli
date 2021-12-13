package file

import (
	"path/filepath"
)

func WhatIsExtension(fileName []string) []string {
	var extensions []string
	for i := 0; i < len(fileName); i++ {
		extensions = append(extensions, filepath.Ext(fileName[i]))
	}
	return extensions
}
