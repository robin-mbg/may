package main
import (
    "fmt"
    "os"
    "github.com/robin-mbg/may/run"
    "github.com/robin-mbg/may/inspect"
)

func main() {
    printSplash()

    if len(os.Args) < 2 {
        fmt.Println("No command provided.")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "run":
        if len(os.Args) < 4 {
            fmt.Println("Command `run` expects path and command as parameters")
            os.Exit(1)
        }

        run.RunCommand(os.Args[2], os.Args[3])
    case "inspect":
        if len(os.Args) < 3 {
            fmt.Println("Command `inspect` expects path as parameter")
            os.Exit(1)
        }

        inspect.RunInspection(os.Args[2])
    case "help":
        fmt.Println("Try running one of these commands:")
        fmt.Println()
        fmt.Println("may run <path> <command>")
        fmt.Println("may inspect <path>")
    default:
        fmt.Println("Unknown command:", os.Args[1])
        os.Exit(1);
    }

    fmt.Println("Thanks for enjoying may")
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
