package formatter

/*
Basic formatter package, just dumps whatever pushed with args
*/
import (
	"fmt"
	"github.com/apsdehal/go-logger/levels"
)

type BasicFormatter struct {
	Formatter
}

func (b *BasicFormatter) Print(level levels.Level, message string, args ...interface{}) string {
	return fmt.Sprintf(message, args...)
}

func Basic() *BasicFormatter {
	return &BasicFormatter{}
}
