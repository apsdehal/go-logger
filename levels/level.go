package levels

/*
Levels define the intensity of output logged.
While DEBUG is the lowest intensity, CRITICAL sits on the other end
*/

import "fmt"

type Level int

const (
	FATAL Level = iota
	CRITICAL
	ERROR
	INFO
	NOTICE
	WARN
	DEBUG
	TRACE
	INHERIT
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

// Returns a proper string to output for colored logging
func colorString(color int) string {
	return fmt.Sprintf("\033[%dm", int(color))
}

var StringToLevels = map[string]Level{
	"CRITICAL":	CRITICAL,
	"FATAL"	  : FATAL,
	"ERROR"	  : ERROR,
	"INFO"	  : INFO,
	"NOTICE"  : NOTICE,
	"WARN"	  : WARN,
	"DEBUG"	  : DEBUG,
	"TRACE"	  : TRACE,
	"INHERIT" : INHERIT,
}

var LevelsToString = map[Level]string{
	CRITICAL:	"CRITICAL",
	FATAL	:   "FATAL",
	ERROR	:   "ERROR",
	INFO	:  	"INFO",
	NOTICE	:	"NOTICE",
	WARN	:   "WARN",
	DEBUG	:   "DEBUG",
	TRACE	:   "TRACE",
	INHERIT	: 	"INHERIT",
}

// Initializes the map of colors
var LevelColors = map[Level]string{
	FATAL	: colorString(Magenta),
	CRITICAL: colorString(Magenta),
	ERROR	: colorString(Red),
	WARN	: colorString(Yellow),
	NOTICE	: colorString(Green),
	DEBUG	: colorString(Cyan),
	INFO 	: colorString(White),
	TRACE	: colorString(White),
	INHERIT : colorString(White),
}