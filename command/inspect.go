package command

import (
	"github.com/robin-mbg/may/util"
	"os"
)

// Inspect takes a repository name parameter and shows which build tool `may` would use to run commands on that repository.
func Inspect(path string) {
	util.Log("Running inspection on " + path)

	preferredExecutor := GetExecutor(path)

	util.LogDebug("Commands on this project would be run using `" + preferredExecutor + "`.")
}

// GetExecutor takes a file system paths and prints which build tool it would
// use to execute commands for that path
func GetExecutor(path string) string {
	var possibleExecutors = []string{}

	if isDockerProject(path) {
		possibleExecutors = append(possibleExecutors, "docker")
	}
	if isGoProject(path) {
		possibleExecutors = append(possibleExecutors, "go")
	}
	if isYarnProject(path) {
		possibleExecutors = append(possibleExecutors, "yarn")
	}
	if isGradleProject(path) {
		possibleExecutors = append(possibleExecutors, path+"/gradlew")
	}
	if isMakefileProject(path) {
		possibleExecutors = append(possibleExecutors, "make")
	}

	if len(possibleExecutors) == 0 {
		util.LogError("No executor could be determined. Consider adding a Makefile.")
		return ""
	}

	preferredExecutor := possibleExecutors[len(possibleExecutors)-1]

	return preferredExecutor
}

func isMakefileProject(path string) bool {
	testPath := path + "/Makefile"

	if exists(testPath) {
		util.LogDebug(path + " can be run using `make`")
		return true
	}
	return false
}

func isGradleProject(path string) bool {
	testPath := path + "/gradlew"

	if exists(testPath) {
		util.LogDebug(path + " can be run using `gradle`")
		return true
	}
	return false
}

func isYarnProject(path string) bool {
	testPath := path + "/yarn.lock"

	if exists(testPath) {
		util.LogDebug(path + " can be run using `yarn`")
		return true
	}
	return false
}

func isGoProject(path string) bool {
	testPath := path + "/go.mod"

	if exists(testPath) {
		util.LogDebug(path + " can be run using `go`")
		return true
	}
	return false
}

func isDockerProject(path string) bool {
	testPath := path + "/Dockerfile"

	if exists(testPath) {
		util.LogDebug(path + " can be run using `docker`")
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
