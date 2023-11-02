package main

import (
	"fmt"
	"os"

	"glox/log"
)

func main() {
	log.Init()

	nargs := len(os.Args)
	if nargs > 2 {
		log.Logger.Error(fmt.Sprintf("usage: %s [file]", os.Args[0]))
		return
	}

	if nargs == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}

func runFile(path string) {

}

func runPrompt() {

}
