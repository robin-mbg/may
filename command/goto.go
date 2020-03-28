package command

import (
	"github.com/robin-mbg/may/find"
	"github.com/robin-mbg/may/util"
	"os"
)

func Go(name string) {
	// Find candidate
	path := find.FindCandidate(name)

	// Change directory
	os.Chdir(path)

	newdir, err := os.Getwd()
	if err == nil {
		util.Log("Successfully changed to " + newdir)
	}
}
