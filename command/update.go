package command

import (
    "github.com/robin-mbg/may/util"
    "github.com/robin-mbg/may/find"
    "os"
)

var (
    repositories []string
)

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
        util.LogDebug("Checking " + repository + " ...")
    }
}

func applyAllUpdates() {
    for _, repository := range repositories {
        util.Log("Pulling into repository " + repository)
        util.RunCommand("git", "pull", repository)
    }
}

func UpdateDefault() {
    util.LogDebug("Assuming `may update check`")
    Update("check")
}