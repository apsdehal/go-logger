package writer

/*
Writers are responsible for logging data from logger instance to an output
New writers can be written by satisfying the given interface
*/

import (
	"github.com/apsdehal/go-logger/formats"
	"github.com/apsdehal/levels"
)


type Writer interface {
	Write(level levels.LogLevel, message string, args ...interface{})
	Format() format.Format
	SetFormat(format.Format) 
}