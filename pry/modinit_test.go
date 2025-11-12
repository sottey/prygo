package pry

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEnsureTempModuleAddsReplaceFromEnv(t *testing.T) {
	root := findLocalModuleDir("github.com/sottey/prygo")
	if root == "" {
		t.Skip("unable to determine module root")
	}
	dir := t.TempDir()

	t.Setenv("PRYGO_MODULE_ROOT", root)

	if err := ensureTempModule(dir); err != nil {
		t.Fatalf("ensureTempModule failed: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "go.mod"))
	if err != nil {
		t.Fatalf("reading go.mod: %v", err)
	}
	goMod := string(data)
	if !strings.Contains(goMod, "module "+tempModuleName) {
		t.Fatalf("temp go.mod missing module declaration: %q", goMod)
	}
	expectReplace := "replace github.com/sottey/prygo => " + root
	if !strings.Contains(goMod, expectReplace) {
		t.Fatalf("temp go.mod missing replace directive %q.\ncontents:\n%s", expectReplace, goMod)
	}
}
