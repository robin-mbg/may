package util

import (
	"fmt"
	"github.com/gookit/color"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

// Log prints a default log message.
func Log(message string) {
	LogRaw(message)
}

// LogNote prints a slightly elevated log message.
func LogNotice(message string) {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		color.Note.Println(message)
	} else {
		LogRaw(message)
	}
}

// LogDebug prints a debug-level log message.
func LogDebug(message string) {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		color.Light.Println(message)
	} else {
		LogRaw(message)
	}
}

// LogImportant prints an important-level log message.
func LogImportant(message string) {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		color.Warn.Println(message)
	} else {
		LogRaw(message)
	}
}

// LogError prints an error-level log message.
func LogError(message string) {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		color.Error.Println(message)
	} else {
		LogRaw(message)
	}
}

// LogSeparator prints a visual separator.
func LogSeparator() {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		color.Success.Println("─────────────────────")
	}
}

// LogRaw prints a non-colored log message.
func LogRaw(message string) {
	fmt.Println(message)
}
