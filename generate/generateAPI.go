package generate

import (
	"bufio"
	"io"
	"time"

	"log"
	"os"

	"github.com/pterm/pterm"
)

func API(lang string, directory string, projectName string) {
	if directory == "" {
	}
	server, err := os.Open("frameworks/http/gorilla/gorilla-mux.go")
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
		spinnerSuccess, _ := pterm.DefaultSpinner.Start("Iniciando Gorilla Mux API...")
		time.Sleep(time.Second * 2)
		spinnerSuccess.Success()

		errDir := os.Mkdir("./"+directory, 0777)
		if errDir != nil {
			log.Fatal(errorPath)
			pterm.Error.Println("Não pode criar o diretório")
			return
		}
		currentServer, err := os.Create(projectName + "./main.go")

		spinnerSuccessDir, _ := pterm.DefaultSpinner.Start("Criando estrutura de pastas...")
		time.Sleep(time.Second * 2)
		spinnerSuccessDir.UpdateText("Projeto: " + projectName)
		time.Sleep(time.Second * 1)
		spinnerSuccessDir.Success("Projeto construído")
		time.Sleep(time.Second * 1)
		spinnerSuccessDir.Success("OK")

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
	} else {
		pterm.Error.Println("Já existe uma pasta com o mesmo nome do projeto que você está tentando criar")

	}
}
