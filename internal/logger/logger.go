package logger

import (
	"fmt"
	"github.com/briandowns/spinner"
	"time"
)

type Logger struct {
	Verbose bool
	Spinner *spinner.Spinner
}

func NewLogger(verbose bool) Logger {
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)

	return Logger{
		Verbose: verbose,
		Spinner: s,
	}
}

func (l *Logger) Printf(format string, args ...any) {
	fmt.Printf(format, args...)
}

func (l *Logger) PrintfV(format string, args ...any) {
	if !l.Verbose {
		return
	}

	l.Printf(format, args)
}

func (l *Logger) Title(title string) {
	fmt.Printf("██████████ %s ██████████\n", title)
}

func (l *Logger) StartSpinner(spinningMessage, finalMessage string) {
	l.Spinner.Suffix = spinningMessage
	l.Spinner.FinalMSG = finalMessage
	l.Spinner.Start()
}

func (l *Logger) StopSpinner() {
	l.Spinner.Stop()
	fmt.Println()
}
