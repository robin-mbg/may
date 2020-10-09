package find

import (
	"fmt"
	"github.com/karrick/godirwalk"
	"github.com/robin-mbg/may/pkg/util"
	"os"
	"strings"
)

var (
	targetFile          string
	gitRepositoriesList []string
	redList             = sliceToStrMap([]string{"Downloads", "Library", "Pictures", "Videos", "Music", "AppData", "Favorites", "Links", "tmp", "temp", "node_modules", "go", "bin", "snap"})
)

// Candidates takes a filter string and lists all repositories matching that string.
func Candidates(name string, includeAll bool, basePaths []string) []string {
	for _, basePath := range basePaths {
		listGitDirectories(basePath, includeAll)
	}

	if name == "" {
		return gitRepositoriesList
	}

	candidates := []string{}
	for _, path := range gitRepositoriesList {
		if strings.Contains(path, name) {
			candidates = append(candidates, path)
		}
	}

	if len(candidates) < 1 {
		util.LogError("No matching repository found")
		os.Exit(1)
	}

	return candidates
}

func sliceToStrMap(elements []string) map[string]string {
	elementMap := make(map[string]string)
	for _, s := range elements {
		elementMap[s] = s
	}
	return elementMap
}

func listGitDirectories(basepath string, includeAll bool) {
	targetFile = ".git"

	testFile, err := os.Open(basepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer testFile.Close()

	testFileInfo, _ := testFile.Stat()
	if !testFileInfo.IsDir() {
		util.LogError(basepath + " is not a directory!")
		os.Exit(-1)
	}

	file, err := os.Open(basepath)
	if err != nil {
		util.LogError("Failed opening directory.")
	}
	defer file.Close()

	walkError := godirwalk.Walk(basepath, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if !de.IsDir() {
				return godirwalk.SkipThis
			}
			if !includeAll && !isRelevantDirectory(de.Name()) {
				return godirwalk.SkipThis
			}
			if strings.HasSuffix(osPathname, ".git") {
				gitRepositoriesList = append(gitRepositoriesList, strings.TrimSuffix(osPathname, "/.git"))
				return nil
			}
			return nil
		},
		Unsorted: true,
	})

	if walkError != nil {
		util.LogError("Error occurred while searching directory tree.")
	}
}

func isRelevantDirectory(name string) bool {
	if strings.HasPrefix(name, ".") && !strings.HasSuffix(name, ".git") {
		return false
	}
	_, exists := redList[name]
	return !exists
}

func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}
