[![Build Status](https://github.com/axetroy/dvm/workflows/test/badge.svg)](https://github.com/axetroy/dvm/actions)

## dvm

Version Manger for Deno

> Due to Deno's historical reasons, dvm does not support older Deno versions. 
> eg `deno --version` is an invalid command in older versions

Features:

- [x] Cross platform support
- [x] Easy to use
- [x] No runtime dependencies (This is why it is not written in nodejs)
- [x] Zero configuration
- [x] Fully compatible with installed Deno

### Usage

Whether you install Deno now or not will not affect the use of dvm.

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

If you are using Linux/MacOS. you can install it with following command:

```shell
# install latest version
wget -qO- https://raw.githubusercontent.com/axetroy/dvm/master/install.sh | bash
# or install specified version
wget -qO- https://raw.githubusercontent.com/axetroy/dvm/master/install.sh | bash -s v0.1.0
```

Or

Download the executable file for your platform at [release page](https://github.com/axetroy/dvm/releases)

Then set the environment variable.

eg, the executable file is in the `~/bin` directory.

```bash
# ~/.bash_profile
export PATH="$PATH:~/bin"
```

finally, try it out.

```bash
dvm --help
```

### Upgrade

You can re-download the executable and overwrite the original file.

or type the following command to upgrade to the latest version.

```bash
> dvm upgrade
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
make test
```

### License

The [MIT License](LICENSE)