package command

import (
	"github.com/robin-mbg/may/find"
	"github.com/robin-mbg/may/util"
)

// Show lists all git repositories that it finds.
func Show() {
	repositories := find.Repositories()
	util.LogSeparator()

	for _, repository := range repositories {
		util.LogDebug(repository)
	}
}
