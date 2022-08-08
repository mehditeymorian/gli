package logger

import (
	"fmt"
	"github.com/briandowns/spinner"
	"time"
)

type Logger struct {
	Verbose      bool
	Spinner      *spinner.Spinner
	SpinnerTitle string
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

func (l *Logger) Separator() {
	fmt.Println("██████████████████████████████████████████")
}

func (l *Logger) EmptyLine() {
	fmt.Println()
}

func (l *Logger) StartSpinner(title string) {
	l.SpinnerTitle = title
	l.Spinner.Suffix = title
	l.Spinner.Start()
}

func (l *Logger) SetSpinnerMessage(message string) {
	l.Spinner.Suffix = l.SpinnerTitle + " " + message
}

func (l *Logger) StopSpinner(finalMessage string) {
	l.SpinnerTitle = ""
	l.Spinner.FinalMSG = finalMessage
	l.Spinner.Stop()
	fmt.Println()
}
