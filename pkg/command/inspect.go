package command

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// Inspect takes a repository name parameter and shows which build tool `may` would use to run commands on that repository.
func Inspect(paths []string) {
	padding := 5
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)

	fmt.Fprintln(w, "repository\tpreferredExecutor\tpossibleExecutors")
	for _, path := range paths {
		possibleExecutors := getPossibleExecutors(path)
		preferredExecutor := ChoosePreferredExecutor(possibleExecutors)
		joinedPossibleExecutors := strings.Join(possibleExecutors, ",")

		if len(joinedPossibleExecutors) == 0 {
			joinedPossibleExecutors = "N/A"
		}
		if len(preferredExecutor) == 0 {
			preferredExecutor = "N/A"
		}

		fmt.Fprintln(w, path+"\t"+preferredExecutor+"\t"+joinedPossibleExecutors+"\t")
	}
	w.Flush()
}

// GetExecutor takes a path and returns the executor `may` would use to run commands on that repository.
func GetExecutor(path string) string {
	possibleExecutors := getPossibleExecutors(path)
	return ChoosePreferredExecutor(possibleExecutors)
}

// ChoosePreferredExecutor takes a file system paths and prints which build tool it would
// use to execute commands for that path
func ChoosePreferredExecutor(executors []string) string {
	if len(executors) == 0 {
		return ""
	}

	preferredExecutor := executors[len(executors)-1]
	return preferredExecutor
}

func getPossibleExecutors(path string) []string {
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
	if isNpmProject(path) {
		possibleExecutors = append(possibleExecutors, "npm")
	}
	if isGradleProject(path) {
		possibleExecutors = append(possibleExecutors, path+"/gradlew")
	}
	if isMakefileProject(path) {
		possibleExecutors = append(possibleExecutors, "make")
	}

	if len(possibleExecutors) == 0 {
		return []string{}
	}

	return possibleExecutors
}

func isMakefileProject(path string) bool {
	testPath := path + "/Makefile"

	if exists(testPath) {
		return true
	}
	return false
}

func isGradleProject(path string) bool {
	testPath := path + "/gradlew"

	if exists(testPath) {
		return true
	}
	return false
}

func isYarnProject(path string) bool {
	testPath := path + "/yarn.lock"

	if exists(testPath) {
		return true
	}
	return false
}

func isNpmProject(path string) bool {
	testPath := path + "/package-lock.json"

	if exists(testPath) {
		return true
	}
	return false
}

func isGoProject(path string) bool {
	testPath := path + "/go.mod"

	if exists(testPath) {
		return true
	}
	return false
}

func isDockerProject(path string) bool {
	testPath := path + "/Dockerfile"

	if exists(testPath) {
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
