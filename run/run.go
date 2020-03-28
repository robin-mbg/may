package run

import "fmt"

func RunCommand(path string, command string) {
    fmt.Println("Looking for", path)
    fmt.Println("Executing", command)
}

