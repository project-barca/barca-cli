package file

import (
	"bufio"
	"io"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func Move(directory string, ext string) {

	if directory == "" {
	}
	server, err := os.Open("cmd/windows/install_packages.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := server.Close(); err != nil {
			panic(err)
		}
	}()
	r := bufio.NewReader(server)

	_, errorPath := os.Stat("./" + directory)

	if os.IsNotExist(errorPath) {

	} else {

		currentServer, err := os.Create(directory + "./install" + ext)

		if err != nil {
			panic(err)
		}
		defer func() {
			if err := currentServer.Close(); err != nil {
				panic(err)
			}
		}()

		w := bufio.NewWriter(currentServer)

		buf := make([]byte, 1024)
		for {
			n, err := r.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}
			if _, err := w.Write(buf[:n]); err != nil {
				panic(err)
			}
		}

		if err = w.Flush(); err != nil {
			panic(err)
		}
	}

}
