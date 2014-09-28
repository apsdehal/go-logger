# go-logger

A simple go logger for easy logging and debugging of your programs

[Click here to view documentation](http://godoc.org/github.com/apsdehal/go-logger)

# Example

Example [program](examples/example.go) demonstrates how to use the logger.

[![Example Output](examples/example.png)](examples/example.go)

```go
package main

import (
	"go get github.com/apsdehal/go-logger"
)

func main () {
	// Get the instance for logger class, 1 states if we want to use coloring
	log, err := logger.New("test", 1)
	if err != nil {
		panic(err) // Check for error
	}

	// Critically log critical
	log.Critical("This is Critical!")
	// Debug
	log.Debug("This is Debug!")
	// Give the Warning
	log.Warning("This is Warning!")
	// Show the error
	log.Error("This is Error!")
	// Notice
	log.Notice("This is Notice!")
	// Show the info
	log.Info("This is Info!")
}
```

# Install

`go get github.com/apsdehal/go-logger`

Use `go get -u` to update the package.

## Thanks

Thanks goes to all go-loggers out there in github world

## License

The [BSD 3-Clause license][bsd], the same as the [Go language][golang].
[bsd]: http://opensource.org/licenses/BSD-3-Clause
[golang]: http://golang.org/LICENSE
