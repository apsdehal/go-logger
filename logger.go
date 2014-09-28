// Package name declaration
package logger

// Import packages
import (
 	"bytes"
 	"fmt"
 	"io"
 	"log"
)

// Contains color strings for stdout
var colors map[string]int

// Color numbers for stdout
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

// Worker class, Worker is a log object used to log messages and Color specifies
// if colored output is to be produced
type Worker struct {
	Minion *log.Logger
	Color bool
}

// Info class, Contains all the info on what has to logged, time is the current time, Module is the specific module
// For which we are logging, level is the state, importance and type of message logged,
// Message contains the string to be logged, format is the format of string to be passed to sprintf
type Info struct {
	Time time.Time
	Module string
	Level string
	Message string
	format string
}

// Logger class that is an interface to user to log messages, Module is the module for which we are testing
// worker is variable of Worker class that is used in bottom layers to log the message
type Logger {
	Module string
	worker *Worker
}

// Returns a proper string to be outputted for a particular info
func (r *Info) Output() string {
	msg := fmt.Sprintf(r.format, r.Time, r.Level, r.message )
	return msg
}

// Returns an instance of worker class, prefix is the string attached to every log, 
// flag determine the log params, color parameters verifies whether we need colored outputs or not
func NewWorker(prefix string, flag int, color bool) *Worker{
	return &Worker{Minion: log.New(os.Stderr, prefix, flag), Color: color}
}

// Function of Worker class to log a string based on level
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

// Returns a proper string to output for colored logging
func colorString(color color) string {
	return fmt.Sprintf("\033[%dm", int(color))
}

// Initializes the map of colors
func initColors() {
	colors = map[string]int{
		"CRITICAL": colorString(Magenta),
		"ERROR":    colorString(Red),
		"WARNING":  colorString(Yellow),
		"NOTICE":   colorString(Green),
		"DEBUG":    colorString(Cyan)
	}
}

// Returns a new instance of logger class, module is the specific module for which we are logging
// colors defines whether the output is to be colored or not
func (*l Logger) New(module string, color bool) (*Logger, error) {
	initColors()
	newWorker := NewWorker("", 0, color)
	return &Logger{Module: module, worker: newWorker}, nil	
}

// The log commnand is the function available to user to log message, lvl specifies
// the degree of the messagethe user wants to log, message is the info user wants to log
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

// Fatal is just like func l,Critical logger except that it is followed by exit to program
func (*l Logger) Fatal(message string) {
	l.Log("CRITICAL", message)
	os.Exit(1)
}

// Panic is just like func l.Critical except that it is followed by a call to panic
func (*l Logger) Panic(message string) {
	l.Log("CRITICAL", message)
	panic(message)
}

// Critical logs a message at a Critical Level
func (*l Logger) Critical(message string) {
	l.Log("CRITICAL", message)
}

// Error logs a message at Error level
func (*l Logger) Error(message string) {
	l.Log("ERROR", message)
}

// Warning logs a message at Warning level
func (*l Logger) Warning(message string) {
	l.Log("WARNING", message)
}

// Notice logs a message at Notice level
func (*l Logger) Notice(message string) {
	l.Log("NOTICE", message)
}

// Info logs a message at Info level
func (*l Logger) Info(message string) {
	l.Log("INFO", message)
}

// Devug logs a message at Devug level
func (*l Logger) Debug(message string) {
	l.Log("DEBUG", message)
}