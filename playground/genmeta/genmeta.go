package main

import (
	"bytes"
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/sottey/prygo/pry"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	dir, err := parser.ParseDir(token.NewFileSet(), wd, nil, 0)
	if err != nil {
		return err
	}
	dirFiles := make(map[string]map[string]*ast.File, len(dir))
	for name, pkg := range dir {
		dirFiles[name] = pkg.Files
	}
	imp := pry.JSImporter{
		Dir: dirFiles,
	}
	spew.Dump(imp)
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(imp); err != nil {
		return err
	}
	if err := os.WriteFile("meta.go", []byte(
		`package main
import "github.com/sottey/prygo/pry"
func init(){
	pry.InternalSetImports(`+"`"+buf.String()+"`"+`)
}`,
	), 0644); err != nil {
		return err
	}
	return nil
}
