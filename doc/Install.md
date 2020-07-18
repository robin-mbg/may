# Installation

Beware that, currently, only Linux is officially supported. Mac OS X and Microsoft Windows are currently NOT supported.

The most recent version of may is `v1.0.0` and is built for `linux/amd64`, `linux/arm`, `linux/arm64` and `linux/386`. It may (pun intended) work on other architectures as well, but for those you'll currently still have to build it yourself.

Run and enjoy!

## Arch Linux and Arch-based distributions

`may` is available from the Arch User Repository as `may`. Installation via the AUR is easiest with an AUR helper such as `yay`. 

```sh
yay -S may
```

## Other distributions

Download the respective `may` version for your architecture [here](https://github.com/robin-mbg/may/releases). Extract the archive and place the binary in a directory available in `$PATH`.

## Build it yourself

Requirements:
- Go tools: [Installation instructions](https://golang.org/doc/install#install)

1. Clone or download and extract the repository.
2. Have it built and placed in a directory listed in `$PATH` by running the following command inside the repository:

```sh
GOBIN=<your-target-directory> make install-release
```
