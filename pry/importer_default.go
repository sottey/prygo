//go:build !js
// +build !js

package pry

import (
	"go/ast"
	"go/types"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"
)

type packagesImporter struct {
}

func (i packagesImporter) Import(path string) (*types.Package, error) {
	return i.ImportFrom(path, "", 0)
}
func (packagesImporter) ImportFrom(path, dir string, mode types.ImportMode) (*types.Package, error) {
	conf := packages.Config{
		Mode: packages.NeedImports | packages.NeedTypes | packages.NeedDeps | packages.NeedTypesSizes,
		Dir:  dir,
	}
	pkgs, err := packages.Load(&conf, path)
	if err != nil {
		return nil, errors.Wrapf(err, "importing %q", path)
	}
	if len(pkgs) == 0 {
		return nil, errors.Errorf("no packages returned for %q", path)
	}
	pkg := pkgs[0]
	if len(pkg.Errors) > 0 {
		return nil, errors.Wrapf(pkg.Errors[0], "loading %q", path)
	}
	if pkg.Types == nil {
		return nil, errors.Errorf("package %q has no type information", path)
	}
	if pkg.Types.Name() == "" {
		return nil, errors.Errorf("package %q has empty name", path)
	}
	return pkg.Types, nil
}

func getImporter() types.ImporterFrom {
	return packagesImporter{}
}

func (s *Scope) parseDir() (map[string]*ast.File, error) {
	dir := filepath.Dir(s.path)
	if err := ensureTempModule(dir); err != nil {
		return nil, errors.Wrapf(err, "ensuring go.mod in %q", dir)
	}
	conf := packages.Config{
		Fset: s.fset,
		Mode: packages.NeedCompiledGoFiles | packages.NeedSyntax,
		Dir:  dir,
	}
	pkgs, err := packages.Load(&conf, ".")
	if err != nil {
		return nil, errors.Wrapf(err, "parsing dir")
	}
	pkg := pkgs[0]
	files := map[string]*ast.File{}
	for i, name := range pkg.CompiledGoFiles {
		files[name] = pkg.Syntax[i]
	}
	return files, nil
}
