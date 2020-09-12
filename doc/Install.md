# Installation

Beware that, currently, only Linux is officially supported. Mac OS X and Microsoft Windows are currently NOT supported.

The most recent version of may is `v1.1.0` and is built for `linux/amd64`, `linux/arm`, `linux/arm64` and `linux/386`. It may (pun intended) work on other architectures as well, but for those you'll currently still have to build it yourself.

After installation, it is recommended to ensure that all git repositories are configured in a way that does not require interactive input (e.g. for username/password or ssh key passphrases). Additional configuration is also recommended for users running `may` from within Windows Subsystem for Linux.

Run and enjoy!

## Download/Install

### Arch Linux and Arch-based distributions

`may` is available from the Arch User Repository as `may`. Installation via the AUR is easiest with an AUR helper such as `yay`. 

```sh
yay -S may
```

### Other distributions

Download the respective `may` version for your architecture [here](https://github.com/robin-mbg/may/releases). Extract the archive and place the binary in a directory available in `$PATH`.

### Build it yourself

Requirements:
- Go tools: [Installation instructions](https://golang.org/doc/install#install)

1. Clone or download and extract the repository.
2. Have it built and placed in a directory listed in `$PATH` by running the following command inside the repository:

```sh
GOBIN=<your-target-directory> make install-release
```

## Recommended configuration

### Ensure no repository requires interactive input

Running batch commands such as `git pull` across multiple repositories requires them to be configured in a way that no interactive input from the user is needed (e.g. input username and password). Ensuring such configuration is in place is very useful, independent of whether `may` is used or not.

It is recommended to setup access via SSH, an exemplary GitHub guide to do so is provided [here](https://docs.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh). For alternate methods, see also [this StackOverflow answer](https://stackoverflow.com/a/51327559).

### Optimization for Windows Subsystem for Linux

Starting from v1.1, `may` includes the `/mnt/c/Users` folder when searching for git repositories in its default configuration. This may lead to severely degraded performance as file-system operations are passed through both Linux and Windows filesystem layers. It is recommended to set the `MAY_BASEPATH` variable and do so quite strictly to avoid unnecessary file-system load. E.g.:

```sh
export MAY_BASEPATH=/mnt/c/Users/<Your Username>/my-git-repo-folder
```
