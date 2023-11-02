package log

import (
	"io"
	"os"

	"github.com/goinbox/golog"
)

var Logger golog.Logger

type consoleWriter struct {
	io.Writer
}

func (c consoleWriter) Flush() error {
	return nil
}

func (c consoleWriter) Free() {
	return
}

func Init() {
	w := &consoleWriter{
		Writer: os.Stderr,
	}

	Logger = golog.NewSimpleLogger(w, golog.NewSimpleFormater()).
		SetLogLevel(golog.LevelDebug).EnableColor()
}
