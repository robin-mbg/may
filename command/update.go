package command

import (
	"github.com/robin-mbg/may/util"
	"sync"
)

var (
	repositories []string
)

// Status checks the `git status -sb` of all git repositories that it finds.
func Status(requestedRepositories []string) {
	repositories = requestedRepositories
	util.Log("Checking status of all repositories")
	util.LogSeparator()

	var waitGroup sync.WaitGroup
	for _, repository := range repositories {
		waitGroup.Add(1)
		go util.RunAsyncCommand("git", []string{"status", "-sb"}, repository, &waitGroup)
	}
	waitGroup.Wait()
}

// Update calls `git update` on all git repositories that it finds.
func Update(requestedRepositories []string) {
	repositories = requestedRepositories

	util.Log("Pulling all available updates")
	util.LogSeparator()

	var waitGroup sync.WaitGroup
	for _, repository := range repositories {
		waitGroup.Add(1)
		go util.RunAsyncCommand("git", []string{"pull"}, repository, &waitGroup)
	}
	waitGroup.Wait()
}
