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

type Worker struct {
	Minion *log.Logger
	Color bool
}

type Info struct {
	Time time.Time
	module string
	level string
	message string
	format string
}

type Logger {
	Module string
	worker *Worker
}
func (r *Info) Output() string {
	msg := fmt.Sprintf(r.format, r.Time, r.level, r.message )
	return msg
}

func NewWorker(prefix string, flag int, color bool) *Worker{
	return &Worker{Minion: log.New(os.Stderr, prefix, flag), Color: color}
}

func (l *Worker) Log(level Level, calldepth int, info *Info) error {
	if b.Color {
		buf := &bytes.Buffer{}
		buf.Write([]byte(colors[level]))
		buf.Write([]byte(info.Output())
		buf.Write([]byte("\033[0m"))
		return b.Minion.Output(calldepth+1, buf.String())
	} else {
		return b.Minion.Output(calldepth+1, info)
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

func (*l Logger) New(module string, color bool) (*Logger, error) {
	newWorker := NewWorker("", 0, color)
	return &Logger{Module: module, worker: newWorker}, nil	
}