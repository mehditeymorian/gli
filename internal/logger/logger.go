package logger

import (
	"fmt"
	"github.com/briandowns/spinner"
	"strings"
	"time"
)

type Logger struct {
	Verbose      bool
	Spinner      *spinner.Spinner
	SpinnerTitle string
	LogBuffer    *strings.Builder
	Buffering    bool
}

func NewLogger(verbose bool) Logger {
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)

	return Logger{
		Verbose: verbose,
		Spinner: s,
	}
}

func (l *Logger) Buffer() {
	l.Buffering = true
	l.LogBuffer = new(strings.Builder)
}

func (l *Logger) Flush() {
	log := l.LogBuffer.String()
	fmt.Println(log)
	l.LogBuffer = nil
	l.Buffering = false
}

func (l *Logger) Printf(format string, args ...any) {
	if l.Buffering {
		l.LogBuffer.WriteString(fmt.Sprintf(format, args...))
	} else {
		fmt.Printf(format, args...)
	}
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
	l.Buffer()
}

func (l *Logger) SetSpinnerMessage(message string) {
	l.Spinner.Suffix = l.SpinnerTitle + " " + message
}

func (l *Logger) StopSpinner(finalMessage string) {
	l.SpinnerTitle = ""
	l.Spinner.FinalMSG = finalMessage
	l.Spinner.Stop()
	fmt.Println()
	l.Flush()
}
