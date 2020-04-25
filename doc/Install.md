# Installation

Beware that, currently, only Linux is officially supported. OS X support is still experimental, first experiments have shown that, while key functionality is still there, a significant amount of work remains to be done in order to create a smooth experience. Microsoft Windows is currently NOT supported.

May is not yet available via package managers, leaving two options of installation:

- Building it yourself.
- Download a pre-built binary on the [Releases](https://github.com/robin-mbg/may/releases) page.

## Using a pre-built binary

Download a pre-built binary for your architecture [here](https://github.com/robin-mbg/may/releases). Rename the binary to `may` and allow it to be executed using:

```
chmod +x may
```

Place the binary in a directory listed in `$PATH`.

## Build it yourself

Requirements:
- Go tools ((Installation instructions)[https://golang.org/doc/install#install])

1. Clone or download and extract the repository.
2. Build it and have it placed in a directory listed in `$PATH` by running the following command inside the repository:

```
GOBIN=<your-target-directory> make release
```

Run and enjoy!
