package command

import (
	"github.com/robin-mbg/may/util"
)

// Run takes a repository name and a command to be executed in that repository.
// It then determines the build tool with which to execute that command and runs it.
func Run(path string, command string) {
	executor := GetExecutor(path)

	argCommand := []string{command}
	util.LogImportant("Executing " + executor + " " + command + " ...")
	util.LogSeparator()
	util.RunCommand(executor, argCommand, path)
}

// RunSimple is the same as Run, but executes the command without additional arguments.
func RunSimple(path string) {
	executor := GetExecutor(path)

	util.LogImportant("Executing " + executor + " ...")
	util.LogSeparator()
	util.RunCommand(executor, []string{}, path)
}
