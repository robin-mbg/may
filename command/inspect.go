package command

import (
	"github.com/robin-mbg/may/find"
	"github.com/robin-mbg/may/util"
	"os"
)

func Inspect(name string) {
	// Find candidate
	path := find.FindCandidate(name)

	// Run inspection
	util.Log("Running inspection on " + path)

	if isGradleProject(path) {
		util.RunCommand(path + "/gradlew", []string{"tasks"}, path)
		os.Exit(0)
	}

	if isYarnProject(path) {
		util.RunCommand("yarn", []string{"run"}, path)
		os.Exit(0)
	}

	if isGoProject(path) {
		util.RunCommand("go", []string{"help"}, path)
		os.Exit(0)
	}
}

func GetExecutor(path string) string {
	if isGradleProject(path) {
		return path + "/gradlew"
	}
	if isYarnProject(path) {
		return "yarn"
	}
	if isGoProject(path) {
		return "go"
	}

	util.LogError("No executor could be detected for project")
	os.Exit(1)
	return ""
}

func isGradleProject(path string) bool {
	testPath := path + "/gradlew"

	if exists(testPath) {
		util.LogDebug("Specified repository is of type `gradle`")
		return true
	}
	return false
}

func isYarnProject(path string) bool {
	testPath := path + "/yarn.lock"

	if exists(testPath) {
		util.LogDebug("Specified repository is of type `yarn`")
		return true
	}
	return false
}

func isGoProject(path string) bool {
	testPath := path + "/go.mod"

	if exists(testPath) {
		util.LogDebug("Specified repository is of type `golang`")
		return true
	}
	return false
}

// Exists reports whether the named file or directory exists.
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
