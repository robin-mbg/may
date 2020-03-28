package run

import (
    "github.com/robin-mbg/may/util"
    "github.com/robin-mbg/may/find"
    "github.com/robin-mbg/may/inspect"
)

func RunCommand(name string, command string) {
    // Find suitable candidate
    path := find.FindCandidate(name)

    // Extract command executor
    executor := inspect.GetExecutor(path)

    // Execute
    util.LogImportant("Executing " + executor + " " + command + "...")
    util.LogSeperator()
    inspect.RunCommand(executor, command, path)
}

