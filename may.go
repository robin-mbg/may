package main

import (
	"fmt"
	"github.com/robin-mbg/may/command"
	"github.com/robin-mbg/may/util"
	"os"
)

var version = "v1.0.0-alpha"

func main() {
	printSplash()

	if len(os.Args) < 2 {
		util.Log("No command provided.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		if len(os.Args) < 4 {
			util.Log("Command `run` expects `name` and build command as parameters")
			os.Exit(1)
		}

		command.Run(os.Args[2], os.Args[3])
	case "inspect":
		if len(os.Args) < 3 {
			util.Log("Command `inspect` expects `name` as parameter")
			os.Exit(1)
		}

		command.Inspect(os.Args[2])
	case "update":
		command.Update()
	case "status":
		command.Status()
	case "show":
		command.Show()
	case "help":
		util.Log("Try running one of these commands:")
		util.LogDebug("may run <repository-name> <command>")
		util.LogDebug("may inspect <repository-name>")
		util.LogDebug("may show")
		util.LogDebug("may status")
		util.LogDebug("may update")
		os.Exit(0)
	case "version":
		util.Log("You are using version " + version + ".")
		os.Exit(0)
	default:
		if len(os.Args) < 3 {
			util.Log("Unknown command: " + os.Args[1])
			os.Exit(1)
		}

		util.LogDebug("Assuming `may run`")
		command.Run(os.Args[1], os.Args[2])
	}

	util.LogSeperator()
	util.Log("Looks like smooth sailing. Thanks for enjoying may.")
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
