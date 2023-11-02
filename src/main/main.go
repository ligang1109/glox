package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/goinbox/golog"

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
		err := runFile(os.Args[1])
		if err != nil {
			log.Logger.Error("runFile error", golog.ErrorField(err))
		}
	} else {
		runPrompt()
	}
}

func runFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("os.Open error: %w", err)
	}
	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		run(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("scanner.Scan error: %w", err)
	}

	return nil
}

func runPrompt() {

}

func run(line string) {
	fmt.Println(line)
}
