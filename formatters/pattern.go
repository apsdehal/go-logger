package formatter

import (
	"github.com/apsdehal/go-logger/levels"
)

type PatternFormatter struct {
	Formatter
	Pattern string
	Created int64
	Re *regexp.Regexp
}
