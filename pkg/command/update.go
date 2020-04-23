package command

import (
	"github.com/robin-mbg/may/pkg/util"
	"os"
	"sync"
)

var (
	repositories []string
)

// Status checks the `git status -sb` of all git repositories that it is given.
func Status(requestedRepositories []string) {
	repositories = requestedRepositories

	var waitGroup sync.WaitGroup
	for _, repository := range repositories {
		waitGroup.Add(1)
		go util.RunAsyncCommand("git", []string{"status", "-sb"}, repository, &waitGroup)
	}
	waitGroup.Wait()
}

// Update calls `git update` on all git repositories that it is given.
func Update(requestedRepositories []string) {
	repositories = requestedRepositories

	var waitGroup sync.WaitGroup
	for _, repository := range repositories {
		waitGroup.Add(1)

		additionalEnv := "GIT_TERMINAL_PROMPT=0"
		extendedEnv := append(os.Environ(), additionalEnv)

		go util.RunAsyncCommandWithEnvironment("git", []string{"pull"}, repository, &waitGroup, extendedEnv)
	}
	waitGroup.Wait()
}
