English | [中文简体](README_zh-CN.md)

[![Build Status](https://github.com/axetroy/dvm/workflows/test/badge.svg)](https://github.com/axetroy/dvm/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/dvm)](https://goreportcard.com/report/github.com/axetroy/dvm)
![Latest Version](https://img.shields.io/github/v/release/axetroy/dvm.svg)
![License](https://img.shields.io/github/license/axetroy/dvm.svg)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/dvm.svg)

## dvm

`dvm` is a command-line tool to manage Deno versions.

Focus on the simplest way to manage versions.

Features:

- [x] Cross-platform support
- [x] Easy to use
- [x] No runtime dependencies
- [x] Zero configuration
- [x] Fully compatible with installed Deno

| Command                         | Description                                 |
| ------------------------------- | ------------------------------------------- |
| dvm current                     | Display currently activated version of Deno |
| dvm ls                          | List installed versions                     |
| dvm ls-remote                   | List remote versions available for install  |
| dvm install \<version\>\|latest | Download and install specified Deno version |
| dvm uninstall \<version\>       | Uninstall specified Deno version            |
| dvm use \<version\>             | Use specified Deno version                  |
| dvm unused                      | Unused Deno                                 |
| dvm exec \<version\> [commands] | Run Deno command on \<version\>             |
| dvm upgrade [version]           | Upgrade dvm                                 |
| dvm destroy                     | Uninstall dvm                               |

### Usage

Whether you have installed Deno or not will not affect the use of dvm.

```bash
# install
$ dvm install v0.26.0
$ deno --version
deno v0.26.0

# use another version
$ dvm install v0.25.0
$ dvm use v0.25
$ deno --version
deno v0.25.0

# uninstall deno
$ dvm uninstall v0.25.0

# for more command
$ dvm --help
```

### Installation

If you are using Linux/macOS. you can install it with the following command:

```shell
# install latest version
curl -fsSL https://raw.githubusercontent.com/axetroy/dvm/master/install.sh | bash
# or install specified version
curl -fsSL https://raw.githubusercontent.com/axetroy/dvm/master/install.sh | bash -s v0.1.0
```

Or

Download the executable file for your platform at [release page](https://github.com/axetroy/dvm/releases)

Then set the environment variable.

eg, the executable file is in the `~/bin` directory.

```bash
# ~/.bash_profile
export PATH="$PATH:$HOME/bin"
```

then, try it out.

```bash
dvm --help
```

Finally, to use Deno correctly, you also need to set environment variables

```bash
# ~/.bash_profile
export PATH="$PATH:$HOME/.deno/bin"
```

### Upgrade

You can re-download the executable and overwrite the original file.

or run the following command to upgrade

```bash
$ dvm upgrade # upgrade to latest
$ dvm upgrade v0.2.0 # Update to specified version
```

### Uninstall

run the following command to uninstall `dvm` or remove `dvm` executable file and `$HOME/.dvm` folder by manual

```shell
$ dvm destroy
```

### Build from source code

Make sure you have `Golang@v1.13.1` installed.

```shell
$ git clone https://github.com/axetroy/dvm.git $GOPATH/src/github.com/axetroy/dvm
$ cd $GOPATH/src/github.com/axetroy/dvm
$ make build
```

### Test

```bash
$ make test
```

### Related

[justjavac/dvm](https://github.com/justjavac/dvm) Node.js implement

[imbsky/dvm](https://github.com/imbsky/dvm) Reason implement

### License

The [MIT License](LICENSE)
