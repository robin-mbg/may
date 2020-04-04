package command

import (
	"github.com/robin-mbg/may/util"
)

// Show lists all git repositories that it finds.
func Show(repositories []string) {
	for _, repository := range repositories {
		util.LogDebug(repository)
	}
}
