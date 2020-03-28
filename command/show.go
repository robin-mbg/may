package command

import (
    "github.com/robin-mbg/may/util"
    "github.com/robin-mbg/may/find"
)

func Show() {
    repositories := find.FindRepositories()
    util.LogSeperator()

    for _, repository := range repositories {
		util.LogDebug(repository)
	}
}

