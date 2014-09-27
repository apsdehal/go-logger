package logger

import (
 	"bytes"
 	"fmt"
 	"io"
 	"log"
)

var colors []string

type color int

const (
	Black = (iota + 30)
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

type Log struct {
	Worker *log.Logger
	Color bool
}

func (l *Logger) Log(level Level, calldepth int, info *Info) error {
	if b.Color {
		buf := &bytes.Buffer{}
		buf.Write([]byte(colors[level]))
		buf.Write([]byte(info.Formatted()))
		buf.Write([]byte("\033[0m"))
		return b.Worker.Output(calldepth+1, buf.String())
	} else {
		return b.Worker.Output(calldepth+1, info.Formatted())
	}
}

func colorString(color color) string {
	return fmt.Sprintf("\033[%dm", int(color))
}

func initColors() {
	colors = []{
		CRITICAL: colorString(Magenta),
		ERROR:    colorString(Red),
		WARNING:  colorString(Yellow),
		NOTICE:   colorString(Green),
		DEBUG:    colorString(Cyan)
	}
}