package find

import (
	"fmt"
	"github.com/robin-mbg/may/util"
	"os"
	"path/filepath"
	"strings"
)

var (
	targetFolder        string
	targetFile          string
	gitRepositoriesList []string
)

// Repositories returns a list of all git repositories that it can find.
func Repositories() []string {
	basepath := getBasePath()
	listGitDirectories(basepath)

	return gitRepositoriesList
}

// Candidates takes a filter string and lists all repositories matching that string.
func Candidates(name string) []string {
	basepath := getBasePath()

	listGitDirectories(basepath)

	if name == "" {
		return gitRepositoriesList
	}

	candidates := []string{}
	for _, v := range gitRepositoriesList {
		if strings.HasSuffix(v, name) {
			candidates = append(candidates, v)
		}
	}

	if len(candidates) < 1 {
		util.LogError("No matching repository found")
		os.Exit(1)
	}

	return candidates
}

func getBasePath() string {
	mayBasePath := os.Getenv("MAY_BASEPATH")
	if len(mayBasePath) > 0 {
		return mayBasePath
	}

	homeDirectory := os.Getenv("HOME")
	if len(homeDirectory) > 0 {
		return homeDirectory
	}

	util.LogError("Could not determine base path. Make sure either $MAY_BASEPATH or $HOME are set.")
	os.Exit(1)

	return ""
}

func listGitDirectories(basepath string) {
	targetFolder = basepath
	targetFile = ".git"

	// sanity check
	testFile, err := os.Open(targetFolder)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer testFile.Close()

	testFileInfo, _ := testFile.Stat()
	if !testFileInfo.IsDir() {
		util.LogError(targetFolder + " is not a directory!")
		os.Exit(-1)
	}

	err = filepath.Walk(targetFolder, findGitRepository)
}

func findGitRepository(path string, fileInfo os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
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
			}
		}
		matched, err := filepath.Match(targetFile, fileInfo.Name())
		if err != nil {
			fmt.Println(err)
		}

		if matched {
			add := absolute
			gitRepositoriesList = append(gitRepositoriesList, strings.TrimSuffix(add, "/.git"))
			testDir.Close()
			return nil
		}
	}
	return nil
}
