package formatter

/*
formatter package define what and how exactly a message is to be logged.
New formatter can be defined once they satisfy the formatter class
*/

import (
	"github.com/apsdehal/go-logger/levels"
)

type Formatter interface {
	Print(level levels.Level, message string, args ...interface{}) string
}

func Default() Formatter {
	return &Basic{}
}