# please-go: please but in go

> Please don't take this too seriously, this is just a silly application I made for fun when I was bored.

## Background

This is a Go port of [please](https://github.com/bitcynth/please) which is a silly application with no good use.

I ported it now as Go 1.16 has added `Setuid` and `Setgid`, so this is now possible.

I have written the Go version to be as close as possible to the C version, but well it is also bad and without almost any error handling.

## Compatibility

Should work on all Linux distros.
Might work on *BSD and macOS.

## Install

- clone the repo (`git clone https://github.com/bitcynth/please-go.git`)
- compile (`go build` or `make build`)
- install (`make install` as root)
- use (`please help`)

## Build requirements

Go 1.16+ is required.

The feature that Go 1.16 introduced that is required is in [CL 210639](https://go-review.googlesource.com/c/go/+/210639/).

Commit: [d1b1145cace8b968307f9311ff611e4bb810710c](https://go.googlesource.com/go/+/d1b1145cace8b968307f9311ff611e4bb810710c).

## License

See the `LICENSE` file in the root of this repository, the text below is a non-binding summary.

please-go is licensed under the BSD 3-Clause license.
