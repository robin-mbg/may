package main
import (
    "fmt"
    "os"
    "github.com/robin-mbg/may/run"
    "github.com/robin-mbg/may/inspect"
    "github.com/robin-mbg/may/util"
)

func main() {
    printSplash()

    if len(os.Args) < 2 {
        util.Log("No command provided.")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "run":
        if len(os.Args) < 4 {
            util.Log("Command `run` expects path and command as parameters")
            os.Exit(1)
        }

        run.RunCommand(os.Args[2], os.Args[3])
    case "inspect":
        if len(os.Args) < 3 {
            util.Log("Command `inspect` expects path as parameter")
            os.Exit(1)
        }

        inspect.RunInspection(os.Args[2])
    case "help":
        util.Log("Try running one of these commands:")
        util.Log("may run <path> <command>")
        util.Log("may inspect <path>")
    default:
        if len(os.Args) < 3 {
            util.Log("Unknown command: " + os.Args[1])
            os.Exit(1);
        }

        util.LogDebug("Assuming `may run`")
        run.RunCommand(os.Args[1], os.Args[2])
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
