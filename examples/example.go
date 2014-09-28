package main

import (
	"../../go-logger"
)

func main () {
	// Get the instance for logger class
	log, err := logger.New("test", 1)
	if err != nil {
		panic(err) // Check for error
	}

	// Critically log critical
	log.Critical("This is Critical!")
	log.Debug("This is Debug!")
	log.Warning("This is Warning!")
	log.Error("This is Error!")
	log.Notice("This is Notice!")
	log.Info("This is Info!")
}