# May CLI

Easily run commands across git repositories all across your system.

## Features

- Run build tools in any repository from one central place
- Switch between repositories
- Update all your git repositories from one central command-line interface

Available top-level commands:

```
may run
may inspect
may go
may update
```

If `may` is not followed by any of these special commands, `may run` is assumed as default.

### Running build tools

Running build tool commands follows the following format. Note that `run` can be omitted as it is the default.

```
may run <repository-name> <command>
```

The build tool with which the command is to be executed is detected automatically based on the content of the repository. 

- If more than one repository is possible based on the name, an error is thrown. 
- If more than one build tool is possible, the one with the highest precedence is used. The list below is sorted in order of precedence.

List of currently supported tools:

- Gradle
- Yarn
- Go

In order to check what kind of build tool commands are available for repository, use `may inspect <name>`.

### Switching between repositories

In order to switch between git repositories, simply run:

```sh
may go <name>
```

### Updating repositories

In order to see for which of your repositories updates are available, run:
```sh
may update check
```

To pull updates for all repositories, run:
```sh
may update apply
```

Simply running `may update` assumes you meant `may update check`.

## FAQ

What happens if I have two repositories of the same name?

- Two repositories of the same name leads to a naming conflict if specified only by that name. Without additional information, `may` cannot extrapolate which of the repositories you actually mean. What you can do is add further information. `may` always checks the suffix of the full path of a repository, which means that you can add the name of the previous folder as well. `backend` could then become `search/backend` or only `ch/backend` if that is already sufficient.
