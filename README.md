# prygo

prygo - an interactive REPL for Go that allows you to drop into your code at any point.

![Tests](https://github.com/sottey/prygo/actions/workflows/unittest.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/sottey/prygo/pry?status.svg)](https://godoc.org/github.com/sottey/prygo/pry)

![prygo](https://i.imgur.com/yr1BEsK.png)

Example

![prygo Animated Example](https://i.imgur.com/H8hFzPV.gif)
![prygo Example](https://i.imgur.com/0rmwVY7.png)



## Usage

Install prygo
```bash
go get github.com/sottey/prygo
go install -i github.com/sottey/prygo

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


## How does it work?
prygo is built using a combination of meta programming as well as a massive amount of reflection. When you invoke the prygo command it looks at the Go files in the mentioned directories (or the current in cases such as `prygo build`) and processes them. Since Go is a compiled language there's no way to dynamically get in scope variables, and even if there was, unused imports would be automatically removed for optimization purposes. Thus, prygo has to find every instance of `pry.Pry()` and inject a large blob of code that contains references to all in scope variables and functions as well as those of the imported packages. When doing this it makes a copy of your file to `.<filename>.gopry` and modifies the `<filename>.go` then passes the command arguments to the standard `go` command. Once the command exits, it restores the files.

If the program unexpectedly fails there is a custom command `prygo restore [files]` that will move the files back. An alternative is to just remove the `pry.Apply(...)` line.

## Inspiration

prygo is greatly inspired by [Pry REPL](http://pryrepl.org) for Ruby.

## License

prygo is licensed under the MIT license.

Made by [Tristan Rice](https://fn.lc).
