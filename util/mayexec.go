package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
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

func checkExecutableExists(executable string) {
	_, err := exec.LookPath(executable)
	if err != nil {
		LogError("Required executable " + executable + " could not be found.")
		os.Exit(1)
	}
}
