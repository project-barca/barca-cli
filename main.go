package main

import (
	"barca-cli/cli"
	"runtime"
)

func main() {

	if runtime.GOOS == "windows" {
		cli.GenerateWin()
	}

}
