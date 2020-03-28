# May CLI

Easily run commands across git repositories all across your system.

## Features

- Update all your git repositories from one central command-line interface
- Run build tools in any repository from one central place

Available top-level commands:

```
may run
may update
may inspect
```

If `may` is not followed by any of these special commands, `may run` is assumed per default.

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

