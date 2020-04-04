package command

import (
	"github.com/robin-mbg/may/find"
	"github.com/robin-mbg/may/util"
)

var (
	repositories []string
)

// Status checks the `git status -sb` of all git repositories that it finds.
func Status() {
	repositories = find.Repositories()
	util.Log("Checking status of all repositories")
	util.LogSeparator()

	for _, repository := range repositories {
		util.Log("Status of " + repository)
		util.RunCommand("git", []string{"status", "-sb"}, repository)
	}
}

// Update calls `git update` on all git repositories that it finds.
func Update() {
	repositories = find.Repositories()

	util.Log("Pulling all available updates")
	util.LogSeparator()

	for _, repository := range repositories {
		util.Log("Pulling into repository " + repository)
		util.RunCommand("git", []string{"pull"}, repository)
	}
}
