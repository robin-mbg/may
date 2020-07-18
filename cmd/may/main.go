package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/robin-mbg/may/pkg/command"
	"github.com/robin-mbg/may/pkg/find"
	"github.com/robin-mbg/may/pkg/util"
	flag "github.com/spf13/pflag"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

var version = "v1.1.0"
var defaultOperation = "show"

func main() {
	startTime := time.Now()
	flag.ErrHelp = errors.New("Please submit any suggestions or issues on https://github.com/robin-mbg/may")

	// Operations
	var operationUpdate = flag.BoolP("Update", "U", false, "(Operation) Trigger git pull operation.")
	var operationStatus = flag.BoolP("Status", "S", false, "(Operation) Trigger git status operation.")
	var operationRun = flag.BoolP("Run", "R", false, "(Operation) Trigger build tool in found repositories.")
	var operationInspect = flag.BoolP("Inspect", "I", false, "(Operation) Show which build tool would be used for given repositories.")
	var operationVersion = flag.BoolP("Version", "V", false, "(Operation) Print currently used version.")

	// Options
	var verbosity = flag.BoolP("verbose", "v", false, "Increase output verbosity.")
	var filter = flag.StringP("filter", "f", "", "Filter repository set according to this criterion.")
	var includeAll = flag.BoolP("all", "a", false, "Search all directories, including dotfiles.")

	flag.Parse()

	// Set and validate operation
	var operations = []string{}

	if *operationUpdate {
		operations = append(operations, "update")
	}
	if *operationStatus {
		operations = append(operations, "status")
	}
	if *operationRun {
		operations = append(operations, "run")
	}
	if *operationInspect {
		operations = append(operations, "inspect")
	}
	if *operationVersion {
		operations = append(operations, "version")
	}
	if len(operations) > 1 {
		util.LogError("You cannot specify more than one operation. See may --help to check which are operations.")
		os.Exit(1)
	}

	chosenOperation := defaultOperation
	if len(operations) > 0 {
		chosenOperation = operations[0]
	}

	// Execute chosen operation ------------------------
	if *verbosity {
		printSplash()
	}
	isRuntimeSupported(*verbosity)

	repositories := []string{}

	pipedInput := readStdIn()
	if len(pipedInput) > 0 {
		repositories = pipedInput
	} else if !isHelperOperation(chosenOperation) {
		repositories = find.Candidates(*filter, *includeAll)
	}

	runOperation(chosenOperation, repositories)

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
		command.Inspect(repositories)
	case "version":
		util.Log(version)
	default:
		util.LogError("An error occurred when deciding the chosen operation.")
		os.Exit(1)
	}
}

func isHelperOperation(operation string) bool {
	if operation == "version" {
		return true
	}
	return false
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

func isRuntimeSupported(verbosity bool) {
	switch os := runtime.GOOS; os {
	case "darwin":
		if verbosity {
			util.Log("OS X support is still experimental. Beware that significant parts of may's functionality may not work as intended.")
		}
	case "linux":
		if verbosity {
			util.Log("Linux is an officially supported runtime. If you find any issues, please submit an issue on Github.")
		}
	default:
		util.LogError("Runtime currently not supported. Only Linux is officially supported, OS X support is still experimental.")
	}
}

func readStdIn() []string {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return []string{}
	}

	output := []string{}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}
		input = strings.TrimSuffix(input, "\n")

		output = append(output, input)
	}

	return output
}
