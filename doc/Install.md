# Installation

Beware that, currently, only Linux is officially supported. Mac OS X and Microsoft Windows are currently NOT supported.

The most current version of may is `v1.0.0`. 

## Arch Linux and Arch-based distributions

`may` is available from the Arch User Repository as `may`. Installation via the AUR is easiest with an AUR helper such as `yay`. 

```sh
yay -S may
```

## Other distributions

Download a pre-built binary for your architecture [here](https://github.com/robin-mbg/may/releases). Rename the binary to `may` and allow it to be executed using:

```
chmod +x may
```

Place the binary in a directory listed in `$PATH`.

## Build it yourself

Requirements:
- Go tools: [Installation instructions](https://golang.org/doc/install#install)

1. Clone or download and extract the repository.
2. Build it and have it placed in a directory listed in `$PATH` by running the following command inside the repository:

```
GOBIN=<your-target-directory> make release
```

Run and enjoy!
