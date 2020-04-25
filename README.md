# May CLI
[![GoReport](https://goreportcard.com/badge/github.com/robin-mbg/may)](https://goreportcard.com/report/github.com/robin-mbg/may) ![GitHub All Releases](https://img.shields.io/github/downloads/robin-mbg/may/total) ![GitHub](https://img.shields.io/github/license/robin-mbg/may)

pacman-inspired tool to easily list and manage git repositories all across your system. It was created to ease the pain of managing an ever-increasing number of repositories and allow easy access to any of them, no matter the directory your shell happens to be in. `may` is also designed to be easily integrated into shell scripts dealing with repositories.

## Installation

For installation instruction see `/doc/Install.md` or click [here](doc/Install.md).

## Features

- Get a list of all git repositories on your system (plain `may`)
- Get a filtered list of your repositories (`may -f <subpath>`)
- Run build tools on repository selection from one central place (`may -R` / `may -Rf <subpath>`)
- View all repositories and check their status (`may -S`)
- Update all your git repositories from one central command-line interface (`may -U`)
- Use `may` in pipes to receive or send git repository lists (e.g. `may | fzf | may -U`)

Available operations:

```
may                 # Lists all repositories available in your home directory
may -R              # `Run`: Runs a build-tool command on an auto-detected build tool. Takes a required positional argument at the end.
may -I              # `Inspect`: Shows commands available for a repository from auto-detected build tool
may -S              # `Status`: Runs `git status` on all repositories may can find
may -U              # `Update`: Allows updating (aka pulling) all repositories at once

may -V              # `Version`: (Helper) Prints the version of `may` currently in use
may --help          # `Help`: (Helper) Prints helpful information
```

Options:
```
may -f <subpath>    # Filter repository list by the given subpath
may -v              # Verbose output
may -a              # All directories are searched, including dotfiles
```

Every call to `may` can consist of 0..1 operations and 0..n options. This means that all of the following are permitted: `may`, `may -Uvf subpath`, `may -Iv`, `may -vI`. The following are NOT permitted: `may -IU`, `may --U`.

### Viewing repositories: `may`, `may -S`, `may -I`

In order to view all repositories available in your home directory, simply run:

```sh
may
```
In order to check the state of all those repositories, use:

```sh
may -S
```
Behind the scenes, this asynchronously executes `git status -sb` for all repositories.

To inspect your repositories further and see which build tools may would be able to use when running `may -R`, use:

```sh
may -I
```
The inspection output is also very useful for writing scripts based on this information.

### Updating repositories: `may -U`

To pull updates for all repositories, run:
```sh
may -U
```
To, for example, only update `vim` plugins, the following is a handy variant:
```sh
may -Uaf ".vim"
```

### Running in repositories: `may -R`

This is very useful to execute build and run commands on multiple or distant repositories. It is recommended to only use this command in combination with a strict filter (`-f <subpath>`) as, e.g., building a large number of repositories can be an extremely lengthy task.

```sh
may -Rf <repositorypath> <task>
```

The `<task>` and other following parameters are forwarded to an auto-selected build tool. In order to find out which build tool would be selected, use `may -I`. It is recommended to add a `Makefile` to your project to concretely define available tasks.

Currently supported are the following tools:

- make
- gradle
- npm
- yarn

Limited/Beta support is available for these tools:

- Docker
- go

Note that by adding a `Makefile`, any command/build tool is easily supported.

### Help: `may --help`, `may -V`

Two helper commands are available, `may --help` to view a short list of generally available commands and `may -V` to check which version you are currently using.

## Customizing

Per default, `may` uses the entire content in `$HOME` for its find operations. You can change this behaviour by setting `MAY_BASEPATH` to your chosen path.

## Scripting

`may` can read paths from `stdin` that then replace the default list of available git repositories. Its output can also easily be used in standard commands such as `awk`, `grep`, ...

The following is a very simple, `fzf`-based example to have a multi-selet of git repositories to pull.

```sh
may | fzf -m | may -U
```

The following builds all `gradle` projects on your system, making use of the responsible `gradlew` in each repository.

```sh
may -I | grep "gradle" | awk '{ print $1 }' | may -R build
```
