package pry

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
)

const (
	tempModuleName = "goprytmp"
	modulePath     = "github.com/sottey/prygo"
)

var (
	moduleInfo     moduleMetadata
	moduleInfoOnce sync.Once
)

type moduleMetadata struct {
	path        string
	version     string
	replacePath string
	goVersion   string
	toolchain   string
}

// ensureTempModule writes a go.mod into dir when neither the directory nor any
// of its parents already belongs to a module tree. This lets Go tooling operate
// inside generated scratch dirs while keeping existing modules untouched.
func ensureTempModule(dir string) error {
	hasMod, err := hasAncestorGoMod(dir)
	if err != nil {
		return err
	}
	if hasMod {
		return nil
	}

	meta := getModuleMetadata()
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "module %s\n\n", tempModuleName)
	goVer, toolchain := moduleLanguageSettings(meta)
	fmt.Fprintf(&buf, "go %s\n", goVer)
	if toolchain != "" {
		fmt.Fprintf(&buf, "toolchain %s\n", toolchain)
	}
	fmt.Fprintln(&buf)
	if meta.path != "" && meta.version != "" {
		fmt.Fprintf(&buf, "\nrequire %s %s\n", meta.path, meta.version)
		if meta.replacePath != "" {
			fmt.Fprintf(&buf, "\nreplace %s => %s\n", meta.path, meta.replacePath)
		}
	}

	modPath := filepath.Join(dir, "go.mod")
	return os.WriteFile(modPath, buf.Bytes(), 0644)
}

func hasAncestorGoMod(dir string) (bool, error) {
	current := dir
	for {
		modPath := filepath.Join(current, "go.mod")
		_, err := os.Stat(modPath)
		if err == nil {
			return true, nil
		}
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return false, err
		}

		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}
	return false, nil
}

func getModuleMetadata() moduleMetadata {
	moduleInfoOnce.Do(func() {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			return
		}
		moduleInfo.path = modulePath
		moduleInfo.version = info.Main.Version
		if moduleInfo.version == "" || moduleInfo.version == "(devel)" {
			moduleInfo.version = "v0.0.0"
		}
		if dir := findLocalModuleDir(modulePath); dir != "" {
			moduleInfo.replacePath = dir
		}
		moduleInfo.goVersion, moduleInfo.toolchain = detectModuleGoSettings(moduleInfo.replacePath)
	})
	return moduleInfo
}

func moduleLanguageSettings(meta moduleMetadata) (goVersion, toolchain string) {
	goVersion, toolchain = meta.goVersion, meta.toolchain
	if goVersion == "" {
		goVersion, toolchain = runtimeGoSettings()
	}
	if goVersion == "" {
		goVersion = "1.22"
	}
	return
}

func detectModuleGoSettings(dir string) (goVersion, toolchain string) {
	if dir == "" {
		return runtimeGoSettings()
	}
	path := filepath.Join(dir, "go.mod")
	goVersion, toolchain = parseGoModDirectives(path)
	if goVersion == "" {
		return runtimeGoSettings()
	}
	return
}

func parseGoModDirectives(path string) (goVersion, toolchain string) {
	file, err := os.Open(path)
	if err != nil {
		return "", ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		if strings.HasPrefix(line, "module ") {
			continue
		}
		if strings.HasPrefix(line, "go ") && goVersion == "" {
			goVersion = strings.TrimSpace(strings.TrimPrefix(line, "go "))
			continue
		}
		if strings.HasPrefix(line, "toolchain ") && toolchain == "" {
			toolchain = strings.TrimSpace(strings.TrimPrefix(line, "toolchain "))
			continue
		}
		if goVersion != "" && toolchain != "" {
			break
		}
	}
	return goVersion, toolchain
}

func runtimeGoSettings() (goVersion, toolchain string) {
	version := strings.TrimPrefix(runtime.Version(), "go")
	if version == "" {
		return "", ""
	}
	clean := trimVersionSuffix(version)
	if clean == "" {
		return "", ""
	}
	goVersion = clean
	toolchain = "go" + clean
	return
}

func trimVersionSuffix(version string) string {
	for i, r := range version {
		if r != '.' && (r < '0' || r > '9') {
			return version[:i]
		}
	}
	return version
}

func findLocalModuleDir(modulePath string) string {
	if modulePath == "" {
		return ""
	}
	if envRoot := os.Getenv("PRYGO_MODULE_ROOT"); envRoot != "" {
		if matchesModule(filepath.Join(envRoot, "go.mod"), modulePath) {
			return envRoot
		}
	}
	for _, start := range []string{cwdOrEmpty(), executableDir()} {
		if start == "" {
			continue
		}
		if dir := searchModuleDir(start, modulePath); dir != "" {
			return dir
		}
	}
	return ""
}

func cwdOrEmpty() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir
}

func executableDir() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(exe)
}

func searchModuleDir(start, modulePath string) string {
	current := start
	for {
		modPath := filepath.Join(current, "go.mod")
		if matchesModule(modPath, modulePath) {
			return current
		}
		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}
	return ""
}

func matchesModule(modPath, modulePath string) bool {
	file, err := os.Open(modPath)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		if strings.HasPrefix(line, "module ") {
			name := strings.TrimSpace(strings.TrimPrefix(line, "module "))
			return name == modulePath
		}
		break
	}
	return false
}
