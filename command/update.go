package command

import (
	"github.com/robin-mbg/may/find"
	"github.com/robin-mbg/may/util"
	"os"
)

var (
	repositories []string
)

func UpdateDefault() {
	util.LogDebug("Assuming `may update apply`")
	Update("apply")
}

func Update(modifier string) {
	repositories = find.FindRepositories()

	switch modifier {
	case "check":
		util.Log("Checking for updates")
		util.LogSeperator()

		checkForUpdates()
	case "apply":
		util.Log("Pulling all available updates")
		util.LogSeperator()

		applyAllUpdates()
	default:
		util.LogError("Unknown modifier " + modifier)
		os.Exit(1)
	}
}

func checkForUpdates() {
	for _, repository := range repositories {
        util.Log("Status of " + repository)
        util.RunCommand("git", []string{"status", "-sb"}, repository)
	}

	util.LogError("Feature not implemented yet")
}

func applyAllUpdates() {
	for _, repository := range repositories {
		util.Log("Pulling into repository " + repository)
		util.RunCommand("git", []string{"pull"}, repository)
	}
}
