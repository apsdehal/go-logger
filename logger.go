// Package name declaration
package logger

// Import packages
import (
 	"bytes"
 	"fmt"
 	"log"
 	"time"
 	"os"
 	"sync/atomic"
)

var (
	// Map for te various codes of colors
	colors map[string]string
	
	// Contains color strings for stdout
	sequenceNo uint64
)
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
	Color int
}

// Info class, Contains all the info on what has to logged, time is the current time, Module is the specific module
// For which we are logging, level is the state, importance and type of message logged,
// Message contains the string to be logged, format is the format of string to be passed to sprintf
type Info struct {
	Id uint64
	Time string
	Module string
	Level string
	Message string
	format string
}

// Logger class that is an interface to user to log messages, Module is the module for which we are testing
// worker is variable of Worker class that is used in bottom layers to log the message
type Logger struct {
	Module string
	worker *Worker
}

// Returns a proper string to be outputted for a particular info
func (r *Info) Output() string {
	msg := fmt.Sprintf(r.format, r.Id, r.Time, r.Level, r.Message )
	return msg
}

// Returns an instance of worker class, prefix is the string attached to every log, 
// flag determine the log params, color parameters verifies whether we need colored outputs or not
func NewWorker(prefix string, flag int, color int) *Worker{
	return &Worker{Minion: log.New(os.Stderr, prefix, flag), Color: color}
}

// Function of Worker class to log a string based on level
func (w *Worker) Log(level string, calldepth int, info *Info) error {
	if w.Color != 0 {
		buf := &bytes.Buffer{}
		buf.Write([]byte(colors[level]))
		buf.Write([]byte(info.Output()))
		buf.Write([]byte("\033[0m"))
		return w.Minion.Output(calldepth+1, buf.String())
	} else {
		return w.Minion.Output(calldepth+1, info.Output())
	}
}

// Returns a proper string to output for colored logging
func colorString(color int) string {
	return fmt.Sprintf("\033[%dm", int(color))
}

// Initializes the map of colors
func initColors() {
	colors = map[string]string{
		"CRITICAL": colorString(Magenta),
		"ERROR":    colorString(Red),
		"WARNING":  colorString(Yellow),
		"NOTICE":   colorString(Green),
		"DEBUG":    colorString(Cyan),
		"INFO" : colorString(White),
	}
}

// Returns a new instance of logger class, module is the specific module for which we are logging
// , color defines whether the output is to be colored or not
func New(module string, color int) (*Logger, error) {
	initColors()
	newWorker := NewWorker("", 0, color)
	return &Logger{Module: module, worker: newWorker}, nil	
}

// The log commnand is the function available to user to log message, lvl specifies
// the degree of the messagethe user wants to log, message is the info user wants to log
func (l *Logger) Log(lvl string, message string) {
	var formatString string = "#%d %s ▶ %.3s %s"
	info := &Info{
		Id:      atomic.AddUint64(&sequenceNo, 1),
		Time:    time.Now().Format("2006-01-02 15:04:05") ,
		Module:  l.Module,
		Level:   lvl,
		Message: message,
		format:  formatString,
	}
	l.worker.Log(lvl, 2, info)
}

// Fatal is just like func l,Cr.tical logger except that it is followed by exit to program
func (l *Logger) Fatal(message string) {
	l.Log("CRITICAL", message)
	os.Exit(1)
}

// Panic is just like func l.Critical except that it is followed by a call to panic
func (l *Logger) Panic(message string) {
	l.Log("CRITICAL", message)
	panic(message)
}

// Critical logs a message at a Critical Level
func (l *Logger) Critical(message string) {
	l.Log("CRITICAL", message)
}

// Error logs a message at Error level
func (l *Logger) Error(message string) {
	l.Log("ERROR", message)
}

// Warning logs a message at Warning level
func (l *Logger) Warning(message string) {
	l.Log("WARNING", message)
}

// Notice logs a message at Notice level
func (l *Logger) Notice(message string) {
	l.Log("NOTICE", message)
}

// Info logs a message at Info level
func (l *Logger) Info(message string) {
	l.Log("INFO", message)
}

// Debug logs a message at Debug level
func (l *Logger) Debug(message string) {
	l.Log("DEBUG", message)
}