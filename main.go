package main

import (
	"barca-cli/cli"
	"fmt"
	"runtime"
)

func main() {

	if runtime.GOOS == "windows" {
		cli.GenerateWin()
	}

}
