# dsha256

> Double-sha256 hash utility

Performs a double-sha256 hash on STDIN.

## Table of Contents
- [Install](#install)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Install

```sh
go get -u github.com/nmarley/dsha256
```

## Build

You can use the standard build process:

```sh
go build
```

... but by using [TinyGo](https://tinygo.org) + strip, the binary can be made
much smaller:

```sh
tinygo build
llvm-strip dsha256
```

A Makefile is included for convenience:

```sh
make
```

## Usage

Example:

```sh
echo 0100000000000000000000000000000000000000000000000000000000000000000000003ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a29ab5f49ffff001d1dac2b7c | xxd -ps -r | dsha256

6fe28c0ab6f1b372c1a6a246ae63f74f931e8365e15a089c68d6190000000000
```

## Contributing

Feel free to dive in! [Open an issue](https://github.com/nmarley/dsha256/issues/new) or submit PRs.

## License

[ISC](LICENSE)
