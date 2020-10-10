package command

import (
	"github.com/robin-mbg/may/pkg/util"
	"strings"
	"sync"
)

// MultiRun takes a list of repositories and a command to be executed in each of them.
func MultiRun(paths []string, fullCommand string, comment string, silent bool) {
	var waitGroup sync.WaitGroup
	for _, path := range paths {
		waitGroup.Add(1)
		var commentOutput = comment
		if len(comment) <= 0 {
			commentOutput = fullCommand
		}

		var headline = ""
		if !silent {
			headline = path + ": " + commentOutput
		}

		executor := strings.Fields(fullCommand)[0]
		argCommand := strings.Fields(fullCommand)[1:]

		go util.RunAsyncCommand(executor, argCommand, path, &waitGroup, headline)
	}

	waitGroup.Wait()
}
