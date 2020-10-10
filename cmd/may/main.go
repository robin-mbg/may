package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/robin-mbg/may/pkg/command"
	"github.com/robin-mbg/may/pkg/find"
	"github.com/robin-mbg/may/pkg/util"
	flag "github.com/spf13/pflag"
	"github.com/stretchr/stew/slice"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

var version = "v1.1.1"
var defaultOperation = "show"
var allowedOperations = []string{"update", "fetch", "pull", "push", "status", "run", "version", "inspect", "show", "help"}

func main() {
	startTime := time.Now()
	flag.ErrHelp = errors.New("Please submit any suggestions or issues on https://github.com/robin-mbg/may")

	// Operations
	var chosenOperation = defaultOperation
	var allArgsWithoutProgram = os.Args[1:]
	if len(allArgsWithoutProgram) > 0 && !strings.HasPrefix(allArgsWithoutProgram[0], "-") {
		chosenOperation = allArgsWithoutProgram[0]
	}

	// Options
	var verbosity = flag.BoolP("verbose", "v", false, "Increased output verbosity.")
	var filter = flag.StringP("filter", "f", "", "Filters repository set according to this criterion.")
	var baseDirectoryArg = flag.StringP("directory", "d", "", "Sets search base directory (default: $HOME + WSL User folder if available).")
	var includeAll = flag.BoolP("all", "a", false, "Searches all directories in base directory, including dotfiles and uncommon directories (e.g. $HOME/Videos).")

	flag.Parse()

	// Validate operation
	if !slice.Contains(allowedOperations, chosenOperation) {
		util.LogError(chosenOperation + " is not a valid command. Use `may help` to list allowed commands.")
		os.Exit(1)
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
		if *baseDirectoryArg == "" {
			repositories = find.Candidates(*filter, *includeAll, getBasePaths())
		} else {
			repositories = find.Candidates(*filter, *includeAll, []string{*baseDirectoryArg})
		}
	}

	runOperation(chosenOperation, repositories)

	if *verbosity {
		executionTime := time.Since(startTime)

		util.LogSeparator()
		util.LogDebug("Execution time: " + executionTime.String())
		util.LogNotice("Looks like smooth sailing. Thanks for enjoying may.")
	}
}

func runOperation(operation string, repositories []string) {
	switch operation {
	case "show":
		command.Show(repositories)
	case "status":
		command.Status(repositories)
	case "fetch":
		util.Log("Not yet implemented.")
	case "pull":
		util.Log("Not yet implemented.")
	case "push":
		util.Log("Generally, batch-pushing multiple repositories does not seem wise and is therefore not implemented. If you really need to do this, use `may run \"git push\"`.")
	case "update":
		command.Update(repositories)
	case "run":
		if len(repositories) > 1 {
			util.LogError("More than one repository found on which command would be executed. This is not yet supported. Try a more specific filter or directory or manually select, e.g. using `may | fzf | may -R`")
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
	case "help":
		util.Log("Not yet implemented.")
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
			util.LogNotice("OS X support is still experimental. Beware that significant parts of may's functionality may not work as intended.")
		}
	case "linux":
		if verbosity {
			util.LogNotice("Linux is an officially supported runtime. If you encounter any problems, please submit an issue on github.com/robin-mbg/may.")
		}
	default:
		util.LogError("Runtime currently not supported. Only Linux is an officially supported runtime, OS X support remains experimental.")
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

func getBasePaths() []string {
	// Setting the basepath explicitly overrides all other options if -d is not set
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
