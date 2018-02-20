
# go-logger

[![Build Status](https://travis-ci.org/apsdehal/go-logger.svg?branch=master)](https://travis-ci.org/apsdehal/go-logger)
[![GoDoc](https://godoc.org/github.com/apsdehal/go-logger?status.svg)](http://godoc.org/github.com/apsdehal/go-logger)

A simple go logger for easy logging and debugging of your programs

# Preview
[![Example Output](examples/example.png)](examples/example.go)

# Formatting
By default all log messages have format that you can see above (on pic).
But you can override the default format and set format that you want.

You can do it for Logger instance (after creating logger) ...
```go
log, _ := logger.New("pkgname", 1)
log.SetFormat(format)
```
... or for package
```go
logger.SetDefaultFormat(format)
```
If you do it for package, all existing loggers will print log messages with format that these used already.
But all newest loggers (which will be created after changing format for package) will use your specified format.

But anyway after this, you can still set format of message for specific Logger instance.

Format of log message must contains verbs that represent some info about current log entry.
Ofc, format can contain not only verbs but also something else (for example text, digits, symbols, etc)

### Format verbs:
You can use the following verbs:
```
%{id}           - means number of current log message
%{module}       - means module name (that you passed to func New())
%{time}			- means current time in format "2006-01-02 15:04:05"
%{time:format}	- means current time in format that you want
					(supports all formats supported by go package "time")
%{level}		- means level name (upper case) of log message ("ERROR", "DEBUG", etc)
%{lvl}			- means first 3 letters of level name (upper case) of log message ("ERR", "DEB", etc)
%{file} 		- means name of file in what you wanna write log
%{filename}		- means the same as %{file}
%{line}			- means line number of file in what you wanna write log
%{message}		- means your log message
```
Non-existent verbs (like ```%{nonex-verb}``` or ```%{}```) will be replaced by an empty string.
Invalid verbs (like ```%{inv-verb```) will be treated as plain text.

# Example

Example [program](examples/example.go) demonstrates how to use the logger.


```go
package main

import (
	"github.com/gjvnq/go-logger"
	"os"
)

func main () {
	// Get the instance for logger class, "test" is the module name, 1 is used to
	// state if we want coloring
	// Third option is optional and is instance of type io.Writer, defaults to os.Stderr
	log, err := logger.New("test", 1, os.Stdout)
	if err != nil {
		panic(err) // Check for error
	}

	// Critically log critical
	log.Critical("This is Critical!")
	log.CriticalF("%+v", err)
	// Debug
	log.Debug("This is Debug!")
	log.DebugF("Here are some numbers: %d %d %f", 10, -3, 3.14)
	// Give the Warning
	log.Warning("This is Warning!")
	log.WarningF("This is Warning!")
	// Show the error
	log.Error("This is Error!")
	log.ErrorF("This is Error!")
	// Notice
	log.Notice("This is Notice!")
	log.NoticeF("%s %s", "This", "is Notice!")
	// Show the info
	log.Info("This is Info!")
	log.InfoF("This is %s!", "Info")

	log.StackAsError("Message before printing stack");

	// Show warning with format
	log.SetFormat("[%{module}] [%{level}] %{message}")
	log.Warning("This is Warning!") // output: "[test] [WARNING] This is Warning!"
	// Also you can set your format as default format for all new loggers
	logger.SetDefaultFormat("%{message}")
	log2, _ := logger.New("pkg", 1, os.Stdout)
	log2.Error("This is Error!") // output: "This is Error!"
}
```

# Install

`go get github.com/apsdehal/go-logger`

Use `go get -u` to update the package.

# Tests

Run `go test logger` to run test on logger
and `go bench` for benchmarks

## Thanks

Thanks goes to all go-loggers out there in github world

## License

The [BSD 3-Clause license][bsd], the same as the [Go language][golang].
[bsd]: http://opensource.org/licenses/BSD-3-Clause
[golang]: http://golang.org/LICENSE
