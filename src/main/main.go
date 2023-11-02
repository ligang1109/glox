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
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("os.Open error: %w", err)
	}
	defer func() {
		_ = f.Close()
	}()

	type errMsg struct {
		msg string
		err error
	}
	var errData []*errMsg

	scanner := bufio.NewScanner(f)
	lineNum := 1
	for scanner.Scan() {
		err = run(scanner.Text())
		if err != nil {
			errData = append(errData, &errMsg{
				msg: fmt.Sprintf("run line %d error", lineNum),
				err: err,
			})
		}

		lineNum++
	}

	for _, item := range errData {
		log.Logger.Error(item.msg, golog.ErrorField(item.err))
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("scanner.Scan error: %w", err)
	}

	return nil
}

func runPrompt() error {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		err := run(scanner.Text())
		if err != nil {
			log.Logger.Error("run error", golog.ErrorField(err))
		}

		fmt.Print("> ")
	}

	err := scanner.Err()
	if err != nil {
		return fmt.Errorf("scanner.Scan error: %w", err)
	}

	return nil
}

func run(line string) error {
	fmt.Println(line)

	return fmt.Errorf("test")
}
