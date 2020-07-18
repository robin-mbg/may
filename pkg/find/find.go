package find

import (
	"fmt"
	"github.com/robin-mbg/may/pkg/util"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	targetFile          string
	gitRepositoriesList []string
	blacklist           = sliceToStrMap([]string{"Downloads", "Pictures", "Videos", "Music", "tmp", "temp", "node_modules", "go", "bin", "snap"})
)

// Candidates takes a filter string and lists all repositories matching that string.
func Candidates(name string, includeAll bool) []string {
	basePaths := getBasePaths()

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

func getBasePaths() []string {
	// Setting the basepath explicitly overrides all other options
	mayBasePath := os.Getenv("MAY_BASEPATH")
	if len(mayBasePath) > 0 {
		return []string{mayBasePath}
	}

	// Default is a combination of $HOME and a possible Windows User folder
	defaultDirectories := []string{}

	homeDirectory := os.Getenv("HOME")
	if len(homeDirectory) > 0 {
		defaultDirectories = append(defaultDirectories, homeDirectory)
	}

	wslMountedDirectory := "/mnt/c/Users"
	_, err := os.Open(wslMountedDirectory)
	if err == nil {
		defaultDirectories = append(defaultDirectories, wslMountedDirectory)
	}

	if len(defaultDirectories) == 0 {
		util.LogError("Could not determine base path. Make sure either $MAY_BASEPATH or $HOME are set.")
		os.Exit(1)
	}

	return defaultDirectories
}

func listGitDirectories(basepath string, includeAll bool) {
	targetFile = ".git"

	// sanity check
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

	// Asynchronously call FileWalkers to search directory tree
	var wg sync.WaitGroup

	list, _ := file.Readdirnames(0)
	for _, name := range list {
		pathWithName := basepath + "/" + name
		shouldBeSearched := isDirectory(pathWithName) && (!includeAll && isRelevantDirectory(name)) || includeAll

		if shouldBeSearched {
			wg.Add(1)
			go func() {
				defer wg.Done()
				filepath.Walk(pathWithName, findGitRepository)
			}()
		}
	}

	wg.Wait()
}

func isRelevantDirectory(name string) bool {
	if strings.HasPrefix(name, ".") {
		return false
	}
	_, exists := blacklist[name]
	return !exists
}

func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func findGitRepository(path string, fileInfo os.FileInfo, err error) error {
	if err != nil {
		return nil
	}

	absolute, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if fileInfo.IsDir() {
		// correct permission to scan folder?
		testDir, err := os.Open(absolute)
		if err != nil {
			if os.IsPermission(err) {
				fmt.Println("No permission to scan ... ", absolute)
				fmt.Println(err)

				testDir.Close()
				return nil
			}
		}

		matched, err := filepath.Match(targetFile, fileInfo.Name())
		if err != nil {
			fmt.Println(err)

			testDir.Close()
			return nil
		}

		if matched {
			pathToAdd := absolute
			gitRepositoriesList = append(gitRepositoriesList, strings.TrimSuffix(pathToAdd, "/.git"))
			testDir.Close()
			return nil
		}
	}
	return nil
}
