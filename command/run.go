package command

import (
    "github.com/robin-mbg/may/util"
    "github.com/robin-mbg/may/find"
)

func Run(name string, command string) {
    // Find suitable candidate
    path := find.FindCandidate(name)

    // Extract command executor
    executor := GetExecutor(path)

    // Execute
    util.LogImportant("Executing " + executor + " " + command + "...")
    util.LogSeperator()
    util.RunCommand(executor, command, path)
}

