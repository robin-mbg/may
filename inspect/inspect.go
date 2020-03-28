package inspect

import (
    "os"
    "fmt"
    "strings"
    "github.com/cheggaaa/pb/v3"
    "path/filepath"
)

var (
    bar          *pb.ProgressBar
    targetFolder string
    targetFile   string
    gitRepositoriesList []string
)

func RunInspection(name string) {
    // List all available repositories
    basepath := os.Getenv("HOME")
    listGitDirectories(basepath)

    candidates := []string{}

    // Find candidates for path
    for _, v := range gitRepositoriesList {
        if strings.HasSuffix(v, name) {
          fmt.Println("Found a match:", v)
          candidates = append(candidates, v)
        }
    }

    if len(candidates) > 1 {
        fmt.Println("Found more than one match")
        os.Exit(1)
    }

    finalCandidate := candidates[0]

    // Run inspection on most promising candidate
    fmt.Println("Running inspection on", finalCandidate)
}

func listGitDirectories(basepath string) {
    count := 10000

    tmpl := `{{ green "Searching for" }} {{string . "path_string" | blue}} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}}`
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
        fmt.Println(targetFolder, " is not a directory!")
        os.Exit(-1)
    }

    err = filepath.Walk(targetFolder, findGitRepository)
    bar.Finish()

    fmt.Println("Found", len(gitRepositoriesList), "git repositories.")
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

