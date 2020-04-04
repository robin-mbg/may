package util

import (
	"github.com/gookit/color"
)

// Log prints a default log message.
func Log(message string) {
	color.Green.Println(message)
}

// LogDebug prints a debug-level log message.
func LogDebug(message string) {
	color.White.Println(message)
}

// LogImportant prints an important-level log message.
func LogImportant(message string) {
	color.Magenta.Println(message)
}

// LogError prints an error-level log message.
func LogError(message string) {
	color.Red.Println(message)
}

// LogSeparator prints a visual separator.
func LogSeparator() {
	color.Green.Println("─────────────────────")
}
