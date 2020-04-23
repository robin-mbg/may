package util

import (
	"fmt"
	"github.com/gookit/color"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

// Log prints a default log message.
func Log(message string) {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		color.Green.Println(message)
	} else {
		LogRaw(message)
	}
}

// LogDebug prints a debug-level log message.
func LogDebug(message string) {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		color.White.Println(message)
	} else {
		LogRaw(message)
	}
}

// LogImportant prints an important-level log message.
func LogImportant(message string) {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		color.Magenta.Println(message)
	} else {
		LogRaw(message)
	}
}

// LogError prints an error-level log message.
func LogError(message string) {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		color.Red.Println(message)
	} else {
		LogRaw(message)
	}
}

// LogSeparator prints a visual separator.
func LogSeparator() {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		color.Green.Println("─────────────────────")
	}
}

// LogRaw prints a non-colored log message.
func LogRaw(message string) {
	fmt.Println(message)
}
