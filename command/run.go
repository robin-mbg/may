package command

import (
    "os"
    "fmt"
    "github.com/robin-mbg/may/util"
    "github.com/robin-mbg/may/find"
)

func Run(name string, command string) {
    path := ""
    if name == "." {
        // Assume current working directory as candidate
        pwd, err := os.Getwd()

        if err != nil {
            util.LogError("Error on extracting current working directory")
            fmt.Println(err)
            os.Exit(0)
        }

        path = pwd
    } else {
        // Find suitable candidate
        path = find.FindCandidate(name)
    }

    // Extract command executor
    executor := GetExecutor(path)

    // Execute
    util.LogImportant("Executing " + executor + " " + command + "...")
    util.LogSeperator()
    util.RunCommand(executor, command, path)
}

