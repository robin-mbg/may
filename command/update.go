package command

import (
	"github.com/robin-mbg/may/find"
	"github.com/robin-mbg/may/util"
)

var (
	repositories []string
)

func Status() {
    repositories = find.FindRepositories()
    util.Log("Checking status of all repositories")
    util.LogSeperator()

    for _, repository := range repositories {
        util.Log("Status of " + repository)
        util.RunCommand("git", []string{"status", "-sb"}, repository)
	}
}

func Update() {
    repositories = find.FindRepositories()

	util.Log("Pulling all available updates")
    util.LogSeperator()

	for _, repository := range repositories {
		util.Log("Pulling into repository " + repository)
		util.RunCommand("git", []string{"pull"}, repository)
	}
}

