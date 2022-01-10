package file

import (
	"fmt"
	"io"
	"os"
)

func Copy(fileName string) {
	// open files r and w
	r, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	w, err := os.Create("clone-" + fileName)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// do the actual work
	n, err := io.Copy(w, r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Copiado %v bytes\n", n)
}
