package logger

import "log"

type Logger struct {
	Verbose bool
}

func (l Logger) Printf(format string, args ...any) {
	log.Printf(format, args)
}

func (l Logger) PrintfV(format string, args ...any) {
	if !l.Verbose {
		return
	}

	l.Printf(format, args)
}
