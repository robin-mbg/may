package main

import (
	"fmt"
	"github.com/robin-mbg/may/command"
	"github.com/robin-mbg/may/find"
	"github.com/robin-mbg/may/util"
	flag "github.com/spf13/pflag"
	"os"
	"time"
)

var version = "v1.0.0-beta"
var defaultOperation = "show"

func main() {
	startTime := time.Now()

	var operation = defaultOperation

	// Operations
	var operationUpdate = flag.BoolP("Update", "U", false, "Trigger git pull operation")
	var operationStatus = flag.BoolP("Status", "S", false, "Trigger git status operation")
	var operationRun = flag.BoolP("Run", "R", false, "Trigger build tool in found repositories")
	var operationInspect = flag.BoolP("Inspect", "I", false, "Show which build tool would be used for given repositories.")

	// Options
	var verbosity = flag.BoolP("verbose", "v", false, "Increase output verbosity")
	var filter = flag.StringP("filter", "f", "", "Filter repository set according to this criterion")

	flag.Parse()

	if *operationUpdate {
		operation = "update"
	}
	if *operationStatus {
		operation = "status"
	}
	if *operationRun {
		operation = "run"
	}
	if *operationInspect {
		operation = "inspect"
	}

	// Execute chosen operation ------------------------
	if *verbosity {
		printSplash()
	}

	repositories := find.Candidates(*filter)
	runOperation(operation, repositories)

	if *verbosity {
		executionTime := time.Since(startTime)

		util.LogSeparator()
		util.LogDebug("Execution time: " + executionTime.String())
		util.Log("Looks like smooth sailing. Thanks for enjoying may.")
	}
}

func runOperation(operation string, repositories []string) {
	switch operation {
	case "show":
		command.Show(repositories)
	case "status":
		command.Status(repositories)
	case "update":
		command.Update(repositories)
	case "run":
		if len(repositories) > 1 {
			util.LogError("More than one repository found on which command would be executed. This is not yet supported.")
			os.Exit(1)
		}
		if len(os.Args) < 4 {
			command.RunSimple(repositories[0])
			return
		}

		command.Run(repositories[0], os.Args[len(os.Args)-1])
	case "inspect":
		for _, repository := range repositories {
			command.Inspect(repository)
		}
	default:
		util.LogError("An error occurred when deciding the chosen operation.")
		os.Exit(1)
	}
}

func printSplash() {
	fmt.Println(",---.    ,---.    ____        ____     __")
	fmt.Println("|    \\  /    |  .'  __ `.     \\   \\   /  /")
	fmt.Println("|  ,  \\/  ,  | /   '  \\  \\     \\  _. /  '")
	fmt.Println("|  |\\_   /|  | |___|  /  |      _( )_ .'")
	fmt.Println("|  _( )_/ |  |    _.-`   |  ___(_ o _)'")
	fmt.Println("| (_ o _) |  | .'   _    | |   |(_,_)'")
	fmt.Println("|  (_,_)  |  | |  _( )_  | |   `-'  /")
	fmt.Println("|  |      |  | \\ (_ o _) /  \\      /")
	fmt.Println("'--'      '--'  '.(_,_).'    `-..-'")
	fmt.Println()
}
