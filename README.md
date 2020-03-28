# May CLI

Easily run commands across git repositories all across your system. It was created to ease the pain of managing an ever-increasing number of repositories and allow easy access to any of them, no matter the directory your shell happens to be in.

## Features

- Run build tools in any repository from one central place
- View all and switch between repositories
- Update all your git repositories from one central command-line interface

Available top-level commands:

```
may run         # Runs a command on an auto-detected build tool
may inspect     # Shows commands available from auto-detected build tool
may show        # Lists all repositories available in your home directory
may go          # Moves to repository of specified name
may update      # Allows updating (aka pulling) all repositories at once
```

If `may` is not followed by any of these special commands, `may run` is assumed as default.

### Running build tools - `may run`, `may inspect`

Running build tool commands follows the following format. Note that `run` can be omitted as it is the default.

```sh
may run <repository-name> <command>
```
If you want to run in your current working directory, you can use `.` as the `<repository-name`, leading to a possible shorthand of `may . build` for `may run <repository-name> build`. Tip: Create an alias in your shell to set `m` to `may`.

The build tool with which the command is to be executed is detected automatically based on the content of the repository. 

- If more than one repository is possible based on the name, an error is thrown. 
- If more than one build tool is possible, the one with the highest precedence is used. The list below is sorted in order of precedence.

List of currently supported tools:

- Gradle
- Yarn
- Go

In order to check what kind of build tool commands are available for repository, use `may inspect <name>`.

### Viewing and switching between repositories - `may show`, `may go`

In order to view all repositories available in your home directory, run:

```sh
may show
```
In order to switch between git repositories, simply run:

```sh
may go <name>
```

### Updating repositories - `may update`

In order to see for which of your repositories updates are available, run:
```sh
may update check
```

To pull updates for all repositories, run:
```sh
may update apply
```

Simply running `may update` assumes you meant `may update apply`.

## FAQ

What happens if I have two repositories of the same name?

- Two repositories of the same name leads to a naming conflict if specified only by that name. Without additional information, `may` cannot extrapolate which of the repositories you actually mean. What you can do is add further information. `may` always checks the suffix of the full path of a repository, which means that you can add the name of the previous folder as well. `backend` could then become `search/backend` or only `ch/backend` if that is already sufficient.
