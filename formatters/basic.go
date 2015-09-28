package formatter

import (
	"fmt"
	"github.com/apsdehal/go-logger/levels"
)

type Basic struct {
	Formatter
}

func (b *Basic) Print(level levels.Level, message string, args ...interface{}) string {
	return fmt.Sprintf(message, args...)
}

