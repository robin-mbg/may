package command

import (
	"github.com/robin-mbg/may/pkg/util"
	"strings"
)

// MultiRunFull takes a list of repositories and a command to be executed in each of them.
func MultiRunFull(paths []string, fullCommand string) {
	for _, path := range paths {
		util.LogImportant(path + ": " + fullCommand)
		RunFull(path, fullCommand)
	}
}

// RunFull takes a repository name and a command to be executed in that repository.
func RunFull(path string, fullCommand string) {
	executor := strings.Fields(fullCommand)[0]

	argCommand := strings.Fields(fullCommand)[1:]
	util.RunCommand(executor, argCommand, path)
}

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
