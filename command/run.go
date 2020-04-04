package command

import (
	"fmt"
	"github.com/robin-mbg/may/find"
	"github.com/robin-mbg/may/util"
	"os"
)

// Run takes a repository name and a command to be executed in that repository.
// It then determines the build tool with which to execute that command and runs it.
func Run(name string, command string) {
	path := ""
	if name == "." {
		// Assume current working directory as candidate
		pwd, err := os.Getwd()

		if err != nil {
			util.LogError("Error on extracting current working directory")
			fmt.Println(err)
			os.Exit(0)
		}

		path = pwd
	} else {
		// Find suitable candidate
		path = find.Candidate(name)
	}

	// Extract command executor
	executor := GetExecutor(path)

	// Execute
	argCommand := []string{command}
	util.LogImportant("Executing " + executor + " " + command + "...")
	util.LogSeparator()
	util.RunCommand(executor, argCommand, path)
}
