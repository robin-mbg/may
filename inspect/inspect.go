package inspect

import (
    "fmt"
    "os"
    "os/exec"
    "github.com/robin-mbg/may/find"
    "github.com/robin-mbg/may/util"
)

func RunInspection(name string) {
    // Find candidate
    path := find.FindCandidate(name)

    // Run inspection
    util.Log("Running inspection on " + path)

    if isGradleProject(path) {
        RunCommand(path + "/gradlew", "tasks", path)
        os.Exit(0)
    }

    if isYarnProject(path) {
        RunCommand("yarn", "run", path)
        os.Exit(0)
    }

    if isGoProject(path) {
        RunCommand("go", "help", path)
        os.Exit(0)
    }
}

func RunCommand(path string, argument string, dir string) {
    cmd := exec.Command(path, argument)
    cmd.Dir = dir

    out, err := cmd.Output()

    if err != nil {
        fmt.Printf("%s", err)
        fmt.Println()
    }

    output := string(out[:])
    fmt.Println(output)
}

func isGradleProject(path string) bool {
    testPath := path + "/gradlew"

    if exists(testPath) {
        util.Log("Specified project is of type `gradle`")
        return true
    }
    return false
}

func isYarnProject(path string) bool {
    testPath := path + "/yarn.lock"

    if exists(testPath) {
        util.Log("Specified project is of type `yarn`")
        return true
    }
    return false
}

func isGoProject(path string) bool {
     testPath := path + "/go.mod"

    if exists(testPath) {
        util.Log("Specified project is of type `golang`")
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
