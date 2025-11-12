# Getting Started with prygo

This guide walks you through installation, basic usage, and your first pry breakpoint.

## Install

```bash
go install github.com/sottey/prygo@latest
```

Ensure `$(go env GOPATH)/bin` (or your `GOBIN`) is on your `PATH`.

## Instrument Your Code

1. Import the pry package:
   ```go
   import "github.com/sottey/prygo/pry"
   ```
2. Drop a breakpoint:
   ```go
   func main() {
       data := compute()
       pry.Pry()
       fmt.Println(data)
   }
   ```

## Run via prygo

Use prygo in place of the `go` command:

```bash
prygo run ./cmd/app
prygo test ./pkg/foo -run TestSpecific
```

When execution reaches `pry.Pry()`, prygo pauses and opens an interactive prompt. Inspect variables, call functions, or mutate state; type `continue` (or press `Ctrl+D`) to resume.

## Standalone REPL

Running prygo with no arguments launches a scratch REPL:

```bash
prygo -i fmt,net/http -e 'fmt.Println("ready")'
```

Flags:

| Flag | Description |
|------|-------------|
| `-i` | Comma-separated packages to pre-import (defaults to `fmt,math`). |
| `-e` | Optional statements to run before entering the REPL. |
| `-generate` | Write the generated pry-injected file to disk instead of running it. |
| `-r` | Automatically revert injected files (true by default). |
| `-d` | Enable debug logs from the generator. |

## Restoring Files

prygo automatically reverts instrumented files when commands finish. If a run crashes (or you abort), you can manually clean up:

```bash
prygo revert ./path/to/package
```

This scans for `.filename.gopry` backups and restores the originals.

## Environment Tips

- prygo mirrors your module’s `go`/`toolchain` directives in temporary workspaces so scratch builds use the same compiler as your project. If prygo can’t find the module root automatically, set `PRYGO_MODULE_ROOT=/absolute/path/to/your/module`.
- For first runs, prygo creates `~/.prygo_history`. Missing history files are fine; prygo falls back to an empty history.

That’s it—drop `pry.Pry()` where you need insight and run your usual `go` command through `prygo`. Next, explore real-world workflows on the [Examples](./Examples.md) page.
