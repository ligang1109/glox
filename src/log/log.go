package log

import (
	"fmt"
	"io"
	"os"

	"github.com/goinbox/golog"
	"github.com/goinbox/gomisc"
)

var Logger golog.Logger

func Init() {
	w := &consoleWriter{
		Writer: os.Stderr,
	}

	Logger = golog.NewSimpleLogger(w, new(flatFormater)).
		SetLogLevel(golog.LevelDebug).EnableColor()
}

type consoleWriter struct {
	io.Writer
}

func (c consoleWriter) Flush() error {
	return nil
}

func (c consoleWriter) Free() {
	return
}

var filterLogFieldKeyMap = map[string]bool{
	"level": true,
	"t":     true,
}

type flatFormater struct {
}

func (f *flatFormater) Format(fields ...*golog.Field) []byte {
	var msg []byte
	for _, field := range fields {
		if !filterLogFieldKeyMap[field.Key] {
			msg = gomisc.AppendBytes(msg, []byte(fmt.Sprintf("%v", field.Value)), []byte("\t"))
		}
	}

	return msg
}
