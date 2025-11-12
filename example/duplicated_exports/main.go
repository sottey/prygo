package main

import (
	"os"
	"path/filepath"

	"github.com/sottey/prygo/pry"
)

func main() {
	a := filepath.Base("/asdf/asdf")
	pry.Pry()
	os.Setenv("foo", "bar")
	_ = a
}
