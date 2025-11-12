# prygo

prygo - an interactive REPL for Go that allows you to drop into your code at any point.

This repository is an actively maintained continuation of Tristan Rice’s original project. The fork keeps pace with modern Go toolchains (Go 1.25+) and focuses on reliability in large modules:

- **Toolchain-aware temp modules** – prygo mirrors your project’s `go`/`toolchain` directives inside scratch workspaces, so `go run` no longer fails with “go mod tidy” or toolchain mismatch errors.
- **Safer import handling** – generated files automatically reference every requested package (including interfaces), preventing “unused import” panics when you drop into the REPL.
- **Improved module discovery** – prygo detects the active module root (or uses `PRYGO_MODULE_ROOT`) so local development uses your checked-out sources without manual replace directives.
- **Up-to-date dependencies & CI** – the module file, GitHub Actions, and tests are refreshed for current Go releases, and failing example binaries are opt-in only.

![Tests](https://github.com/sottey/prygo/actions/workflows/unittest.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/sottey/prygo/pry?status.svg)](https://godoc.org/github.com/sottey/prygo/pry)

![prygo](https://i.imgur.com/yr1BEsK.png)

Example

![prygo Animated Example](https://i.imgur.com/H8hFzPV.gif)
![prygo Example](https://i.imgur.com/0rmwVY7.png)



## Usage

Install prygo
```bash
go install github.com/sottey/prygo@latest
```

Add the pry statement to the code
```go
package main

import "github.com/sottey/prygo/pry"

func main() {
  a := 1
  pry.Pry()
}
```

Run the code as you would normally with the `go` command. prygo is just a wrapper.
```bash
# Run
prygo run readme.go
```

If you want completions to work properly, also install `gocode` if it
is not installed in your system

```bash
go get -u github.com/nsf/gocode
```

More details are available in [examples.md](examples.md)

## How does it work?
prygo is built using a combination of meta programming as well as a massive amount of reflection. When you invoke the prygo command it looks at the Go files in the mentioned directories (or the current in cases such as `prygo build`) and processes them. Since Go is a compiled language there's no way to dynamically get in scope variables, and even if there was, unused imports would be automatically removed for optimization purposes. Thus, prygo has to find every instance of `pry.Pry()` and inject a large blob of code that contains references to all in scope variables and functions as well as those of the imported packages. When doing this it makes a copy of your file to `.<filename>.gopry` and modifies the `<filename>.go` then passes the command arguments to the standard `go` command. Once the command exits, it restores the files.

If the program unexpectedly fails there is a custom command `prygo restore [files]` that will move the files back. An alternative is to just remove the `pry.Apply(...)` line.

## Examples & Tutorials

See [examples.md](examples.md) for a guided tour covering installation, standalone REPL usage, common debugging workflows, and tips for first-time users.

## Inspiration

prygo is greatly inspired by [Pry REPL](http://pryrepl.org) for Ruby.

## License

prygo is licensed under the MIT license.

prygo was refactored and is maintained by [sottey](https://github.com/sottey)

The original go-pry was made by [Tristan Rice](https://fn.lc).
