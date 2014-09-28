package logger

import "testing"

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