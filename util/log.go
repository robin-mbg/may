package util

import (
    "github.com/gookit/color"
)

func Log(message string) {
    color.Green.Println(message)
}

func LogDebug(message string) {
    color.White.Println(message)
}

func LogImportant(message string) {
    color.Magenta.Println(message)
}

func LogError(message string) {
    color.Red.Println(message)
}
