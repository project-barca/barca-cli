package console

import (
	"log"
	"os"
	"os/exec"
)

func StartScript(fileName string) {
	curr_wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("cmd", "start", "/c", curr_wd+"\\"+fileName)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}
