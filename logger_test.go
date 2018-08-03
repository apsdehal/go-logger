package logger

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strings"
	"testing"
	"time"
)

func BenchmarkLoggerLog(b *testing.B) {
	b.StopTimer()
	log, err := New("test", 1)
	if err != nil {
		panic(err)
	}

	var tests = []struct {
		level   LogLevel
		message string
	}{
		{
			CriticalLevel,
			"Critical Logging",
		},
		{
			InfoLevel,
			"Info Logging",
		},
		{
			DebugLevel,
			"Debug logging",
		},
		{
			WarningLevel,
			"Warning logging",
		},
		{
			NoticeLevel,
			"Notice Logging",
		},
		{
			ErrorLevel,
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
	//initColors()
	var tests = []struct {
		level       LogLevel
		color       int
		colorString string
	}{
		{
			CriticalLevel,
			Magenta,
			"\033[35m",
		},
		{
			ErrorLevel,
			Red,
			"\033[31m",
		},
		{
			WarningLevel,
			Yellow,
			"\033[33m",
		},
		{
			NoticeLevel,
			Green,
			"\033[32m",
		},
		{
			DebugLevel,
			Cyan,
			"\033[36m",
		},
		{
			InfoLevel,
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
	var worker *Worker = NewWorker("", 0, 1, os.Stderr, InfoLevel)
	if worker.Minion == nil {
		t.Errorf("Minion was not established")
	}
}

func BenchmarkNewWorker(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		worker := NewWorker("", 0, 1, os.Stderr, InfoLevel)
		if worker == nil {
			panic("Failed to initiate worker")
		}
	}
}

func TestLogger_SetFormat(t *testing.T) {
	var buf bytes.Buffer
	log, err := New("pkgname", 0, &buf)
	if err != nil || log == nil {
		panic(err)
	}
	log.Debug("Test")
	want := time.Now().Format("2006-01-02 15:04:05")
	want = fmt.Sprintf("#1 %s logger_test.go:151 ▶ DEB Test\n", want)
	have := buf.String()
	if have != want {
		t.Errorf("\nWant: %sHave: %s", want, have)
	}
	format :=
		"text123 %{id} " + // text and digits before id
			"!@#$% %{time:Monday, 2006 Jan 01, 15:04:05} " + // symbols before time with spec format
			"a{b %{module} " + // brace with text that should be just text before verb
			"a}b %{filename} " + // brace with text that should be just text before verb
			"%% %{file} " + // percent symbols before verb
			"%{%{line} " + // percent symbol with brace before verb w/o space
			"%{nonex_verb} %{lvl} " + // nonexistent verb berfore real verb
			"%{incorr_verb %{level} " + // incorrect verb before real verb
			"%{} [%{message}]" // empty verb before message in sq brackets
	buf.Reset()
	log.SetFormat(format)
	log.Error("This is Error!")
	now := time.Now()
	want = fmt.Sprintf(
		"text123 2 "+
			"!@#$%% %s "+
			"a{b pkgname "+
			"a}b logger_test.go "+
			"%%%% logger_test.go "+ // it's printf, escaping %, don't forget
			"%%{170 "+
			" ERR "+
			"%%{incorr_verb ERROR "+
			" [This is Error!]\n",
		now.Format("Monday, 2006 Jan 01, 15:04:05"),
	)
	have = buf.String()
	if want != have {
		t.Errorf("\nWant: %sHave: %s", want, have)
		want_len := len(want)
		have_len := len(have)
		min := int(math.Min(float64(want_len), float64(have_len)))
		if want_len != have_len {
			t.Errorf("Diff lens: Want: %d, Have: %d.\n", want_len, have_len)
		}
		for i := 0; i < min; i++ {
			if want[i] != have[i] {
				t.Errorf("Differents starts at %d pos (\"%c\" != \"%c\")\n",
					i, want[i], have[i])
				break
			}
		}
	}
}

func TestSetDefaultFormat(t *testing.T) {
	SetDefaultFormat("%{module} %{lvl} %{message}")
	var buf bytes.Buffer
	log, err := New("pkgname", 0, &buf)
	if err != nil || log == nil {
		panic(err)
	}
	log.Criticalf("Test %d", 123)
	want := "pkgname CRI Test 123\n"
	have := buf.String()
	if want != have {
		t.Errorf("\nWant: %sHave: %s", want, have)
	}
}

func TestLogLevel(t *testing.T) {

	var tests = []struct {
		level   LogLevel
		message string
	}{
		{
			CriticalLevel,
			"Critical Logging",
		},
		{
			ErrorLevel,
			"Error logging",
		},

		{
			WarningLevel,
			"Warning logging",
		},
		{
			NoticeLevel,
			"Notice Logging",
		},
		{
			DebugLevel,
			"Debug logging",
		},
		{
			InfoLevel,
			"Info Logging",
		},
	}

	var buf bytes.Buffer
	log, err := New("pkgname", 0, &buf)
	if err != nil {
		panic(err)
	}

	for i, test := range tests {
		log.SetLogLevel(test.level)

		log.Critical("Log Critical")
		log.Error("Log Error")
		log.Warning("Log Warning")
		log.Notice("Log Notice")
		log.Debug("Log Debug")
		log.Info("Log Info")

		// Count output lines from logger
		count := strings.Count(buf.String(), "\n")
		if i+1 != count {
			t.Error()
		}
		buf.Reset()
	}
}
