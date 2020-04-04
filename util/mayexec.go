package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// RunCommand is a helper function that runs system commands and prints their output to stdout.
func RunCommand(path string, argument []string, dir string) {
	cmd := exec.Command(path, argument...)
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
	//fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}
