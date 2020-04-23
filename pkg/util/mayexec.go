package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

// RunCommand is a helper function that runs system commands and prints their output to stdout.
func RunCommand(executable string, argument []string, dir string) {
	checkExecutableExists(executable)

	cmd := exec.Command(executable, argument...)
	cmd.Dir = dir

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		LogError("Command failed")
		fmt.Println(err)
	}
	outStr, _ := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())

	if len(outStr) < 1 {
		LogDebug("(Command has generated no output)")
	}
}

// RunAsyncCommand is meant for execution as a goroutine and requires a WaitGroup.
// It also does not print directly to stdout, but does so only when the command has terminated.
func RunAsyncCommand(executable string, argument []string, dir string, wg *sync.WaitGroup) {
	env := os.Environ()
	RunAsyncCommandWithEnvironment(executable, argument, dir, wg, env)
}

// RunAsyncCommandWithEnvironment adds functionality for specifying the environment of a process to be run.
// This enables commands like `FOO=BAR mybinary`.
func RunAsyncCommandWithEnvironment(executable string, argument []string, dir string, wg *sync.WaitGroup, env []string) {
	defer wg.Done()
	checkExecutableExists(executable)

	cmd := exec.Command(executable, argument...)
	cmd.Dir = dir
	cmd.Env = env

	out, err := cmd.CombinedOutput()

	Log("Result of " + executable + " " + argument[0] + " in " + dir + ":")
	if err != nil {
		LogError("Command failed")
		fmt.Println(err)
	}

	if len(string(out)) < 1 {
		LogDebug("(Command has generated no output)")
	} else {
		LogDebug(string(out))
	}
}

func checkExecutableExists(executable string) {
	_, err := exec.LookPath(executable)
	if err != nil {
		LogError("Required executable " + executable + " could not be found.")
		os.Exit(1)
	}
}
