package logger

import (
	"testing"
	"os"
)

func BenchmarkLoggerLog(b *testing.B) {
	b.StopTimer()
	log, err := New("test", 1)
	if err != nil {
		panic(err)
	}

	var tests = []struct {
		level string
		message string
	}{
		{
			"CRITICAL",
			"Critical Logging",
		},
		{
			"INFO",
			"Info Logging",
		},
		{
			"DEBUG",
			"Debug logging",
		},
		{
			"WARNING",
			"Warning logging",
		},
		{
			"NOTICE",
			"Notice Logging",
		},
		{
			"ERROR",
			"Error logging",
		},
	}

	b.StartTimer()
	for _, test := range tests {
		for n := 0; n <= b.N; n++ {
			log.Log(test.level, test.message)
		}
	}
}

func BenchmarkLoggerNew(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		log, err := New("test", 1)
		if err != nil && log == nil {
			panic(err)
		}
	}
}

func TestLoggerNew(t *testing.T) {
	log, err := New("test", 1)
	if err != nil {
		panic(err)
	}
	if log.Module != "test" {
		t.Errorf("Unexpected module: %s", log.Module)
	}
}

func TestColorString(t *testing.T) {
	colorCode := colorString(40)
	if colorCode != "\033[40m" {
		t.Errorf("Unexpected string: %s", colorCode)
	}
}

func TestInitColors(t *testing.T) {
	initColors()
	var tests = []struct{
		level string
		color int
		colorString string
	}{
		{
			"CRITICAL",
			 Magenta,
			"\033[35m",
		},
		{
			"ERROR",
			 Red,
			"\033[31m",
		},
		{
			"WARNING",
			 Yellow,
			"\033[33m",
		},
		{
			"NOTICE",
			 Green,
			"\033[32m",
		},
		{
			"DEBUG",
			 Cyan,
			"\033[36m",
		},
		{
			"INFO",
			 White,
			"\033[37m",
		},
	}

	for _, test := range tests {
		if colors[test.level] != test.colorString {
			t.Errorf("Unexpected color string %d", test.color)
		}
	}
}

func TestNewWorker(t *testing.T) {
	var worker *Worker = NewWorker("", 0, 1, os.Stderr)
	if worker.Minion == nil {
		t.Errorf("Minion was not established")
	}
}

func BenchmarkNewWorker(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		worker := NewWorker("", 0, 1, os.Stderr)
		if worker == nil {
			panic("Failed to initiate worker")
		}
	}
}