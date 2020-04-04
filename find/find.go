package find

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/robin-mbg/may/util"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	bar                 *pb.ProgressBar
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

// Candidate takes a repository name and determines the path on which that
// repository is available.
func Candidate(name string) string {
	basepath := getBasePath()
	listGitDirectories(basepath)

	candidates := []string{}

	for _, v := range gitRepositoriesList {
		if strings.HasSuffix(v, name) {
			util.LogDebug("Found a match: " + v)
			candidates = append(candidates, v)
		}
	}

	if len(candidates) < 1 {
		util.LogError("No matching repository found")
		os.Exit(1)
	}

	if len(candidates) > 1 {
		util.LogError("Found more than one match")
		os.Exit(1)
	}

	finalCandidate := candidates[0]

	return finalCandidate
}

func getBasePath() string {
	mayBasePath := os.Getenv("MAY_BASEPATH")
	if len(mayBasePath) > 0 {
		util.LogDebug("Using " + mayBasePath + " for all find operations.")
		return mayBasePath
	}

	homeDirectory := os.Getenv("HOME")
	if len(homeDirectory) > 0 {
		util.LogDebug("Using " + homeDirectory + " for all find operations.")
		return homeDirectory
	}

	util.LogError("Could not determine base path. Make sure either $MAY_BASEPATH or $HOME are set.")
	os.Exit(1)

	return ""
}

func listGitDirectories(basepath string) {
	count := 10000

	tmpl := `{{ green "Finding" }} {{string . "path_string" | blue}} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}}`
	bar = pb.ProgressBarTemplate(tmpl).Start(count)
	bar.Set("path_string", "git repositories")

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
	bar.Finish()

	util.LogDebug("Detected " + strconv.FormatInt(int64(len(gitRepositoriesList)), 10) + " git repositories.")
}

func findGitRepository(path string, fileInfo os.FileInfo, err error) error {
	bar.Increment()
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
