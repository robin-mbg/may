package command

import (
	"github.com/robin-mbg/may/find"
	"github.com/robin-mbg/may/util"
)

func Show() {
	repositories := find.Repositories()
	util.LogSeperator()

	for _, repository := range repositories {
		util.LogDebug(repository)
	}
}
