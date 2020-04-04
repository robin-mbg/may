package command

import (
	"github.com/robin-mbg/may/util"
	"os"
)

// Inspect takes a repository name parameter and shows which build tool `may` would use to run commands on that repository.
// It also calls on that tool to print available commands.
func Inspect(path string) {
	util.Log("Running inspection on " + path)

	var preferredExecutor = ""

	if isDockerProject(path) {
		preferredExecutor = "docker"
	}
	if isGoProject(path) {
		preferredExecutor = "go"
	}
	if isYarnProject(path) {
		preferredExecutor = "yarn"
	}
	if isGradleProject(path) {
		preferredExecutor = "gradle"
	}
	if isMakefileProject(path) {
		preferredExecutor = "make"
	}

	if preferredExecutor == "" {
		util.LogError("No executor could be determined. Consider adding a Makefile.")
		return
	}

	util.LogDebug("Commands on this project would be run using `" + preferredExecutor + "`.")
}

// GetExecutor takes a file system paths and prints which build tool it would
// use to execute commands for that path
func GetExecutor(path string) string {
	if isMakefileProject(path) {
		return "make"
	}
	if isGradleProject(path) {
		return path + "/gradlew"
	}
	if isYarnProject(path) {
		return "yarn"
	}
	if isGoProject(path) {
		return "go"
	}
	if isDockerProject(path) {
		return "docker"
	}

	util.LogError("No executor could be detected for project")
	os.Exit(1)
	return ""
}

func isMakefileProject(path string) bool {
	testPath := path + "/Makefile"

	if exists(testPath) {
		util.LogDebug(path + " has a Makefile")
		return true
	}
	return false
}

func isGradleProject(path string) bool {
	testPath := path + "/gradlew"

	if exists(testPath) {
		util.LogDebug(path + " is of type `gradle`")
		return true
	}
	return false
}

func isYarnProject(path string) bool {
	testPath := path + "/yarn.lock"

	if exists(testPath) {
		util.LogDebug(path + " is of type `yarn`")
		return true
	}
	return false
}

func isGoProject(path string) bool {
	testPath := path + "/go.mod"

	if exists(testPath) {
		util.LogDebug(path + " is of type `golang`")
		return true
	}
	return false
}

func isDockerProject(path string) bool {
	testPath := path + "/Dockerfile"

	if exists(testPath) {
		util.LogDebug(path + " is of type `docker`")
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
