package main

import (
	"os"
	"github.com/apsdehal/go-logger"
)

func main () {
	// Get the instance for logger class
	// Third option is optional and is instance of type io.Writer, defaults to os.Stderr
	log, err := logger.New("test", 1, os.Stdout)
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

	// Show warning with format message
	// Verbs:
	// %{id} - number of msg
	// %{time} - time with format 2006-01-02 15:04:05
	//		%{time:format} - time with specified format
	// %{module} - module name
	// %{file} or %{filename} - filename
	// %{line} - line number in file
	// %{level} or %{lvl} - log level
	//		%{lvl} - print only 3 first letters of log level
	//		for example: DEB, WAR, etc
	//		%{level} - print full name of log level
	//	%{message} - log message
	log.SetFormat("[%{module}] [%{level}] %{message}")
	log.Warning("This is Warning!")
	// Also you can set your format as default format for all new loggers
	logger.SetDefaultFormat("%{message}")
	log2, _ := logger.New("pkg", 1, os.Stdout)
	log2.Error("This is Error!")
}