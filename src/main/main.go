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
		err := runPrompt()
		if err != nil {
			log.Logger.Error("runPrompt error", golog.ErrorField(err))
		}
	}
}

func runFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("os.ReadFile error: %w", err)
	}

	run(string(content))

	return nil
}

func runPrompt() error {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		run(scanner.Text())
		fmt.Print("> ")
	}

	err := scanner.Err()
	if err != nil {
		return fmt.Errorf("scanner.Scan error: %w", err)
	}

	return nil
}

func run(source string) {
	fmt.Println(source)
}
