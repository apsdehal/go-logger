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
	Module string
	Level string
	Message string
	format string
}

type Logger {
	Module string
	worker *Worker
}
func (r *Info) Output() string {
	msg := fmt.Sprintf(r.format, r.Time, r.Level, r.message )
	return msg
}

func NewWorker(prefix string, flag int, color bool) *Worker{
	return &Worker{Minion: log.New(os.Stderr, prefix, flag), Color: color}
}

func (l *Worker) Log(level string, calldepth int, info *Info) error {
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
	colors = map[string]int{
		"CRITICAL": colorString(Magenta),
		"ERROR":    colorString(Red),
		"WARNING":  colorString(Yellow),
		"NOTICE":   colorString(Green),
		"DEBUG":    colorString(Cyan)
	}
}

func (*l Logger) New(module string, color bool) (*Logger, error) {
	initColors()
	newWorker := NewWorker("", 0, color)
	return &Logger{Module: module, worker: newWorker}, nil	
}

func (*l Logger) Log(lvl string, message string) {
	var formatString string = "Some format here"
	info := &Info{
		Time: timeNow(),
		Module: l.Module,
		Level: lvl,
		Message: message,
		format: formatString
	}
	l.worker.log(lvl, 2, info)
}