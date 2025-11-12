# prygo Examples

This guide walks through the most common prygo workflows so that new users can go from “what is this?” to “debugging with prygo” in a few minutes.

## What Is prygo?

prygo is an interactive debugger/REPL for Go programs. You instrument your Go code with `pry.Pry()` statements, run your program through the `prygo` wrapper instead of the raw `go` tool, and prygo rewrites the code on the fly:

1. It replaces each `pry.Pry()` with a call that captures every in‑scope variable plus every imported package symbol.
2. It executes the command you asked for (e.g. `run`, `test`, `build`, or no command for a standalone REPL).
3. When execution reaches `pry.Pry()`, you get an interactive shell where you can evaluate Go expressions, call functions, inspect data, and then continue execution.

The workflow is similar to Ruby’s Pry REPL, but implemented in pure Go by manipulating your source just before compilation.

## Installation

```bash
go install github.com/sottey/prygo@latest
```

Ensure `GOBIN` (or `$(go env GOPATH)/bin`) is on your `PATH`, then run `prygo` instead of `go`.

## Quick Start

1. Import the pry package in your Go file:
   ```go
   import "github.com/sottey/prygo/pry"
   ```
2. Drop a `pry.Pry()` where you want to stop:
   ```go
   func main() {
       total := calculate()
       pry.Pry()
       fmt.Println(total)
   }
   ```
3. Run your program through prygo:
   ```bash
   prygo run ./cmd/app
   ```
4. When execution hits `pry.Pry()`, prygo opens an interactive prompt that has access to every variable in scope plus the exported identifiers from all imported packages.

Type Go expressions (e.g. `total`, `strings.Join(values, ",")`, `continue`) to inspect, mutate, or resume your program.

## Running the Standalone REPL

Invoke prygo with no arguments to start a scratch REPL:

```bash
prygo
```

Flags you can pass to the standalone REPL:

| Flag | Default | Description |
|------|---------|-------------|
| `-i` | `fmt,math` | Comma‑separated list of packages to pre-import. prygo automatically references each package so the generated stub compiles without “unused import” errors. |
| `-e` | `""` | Optional statements to execute before dropping into the REPL. |
| `-generate <path>` | `""` | Emit a pry-injected file at the given path (instead of running it immediately). |
| `-r` | `true` | Revert any modified files when the command exits. Usually keep this true. |
| `-d` | `false` | Enable debug logging from the generator/injector. Helpful when troubleshooting. |

Example session:

```bash
prygo -i fmt,net/http
```

Inside the prompt:

```
[0] prygo> fmt.Println("hello from prygo")
"hello from prygo"
=> 14
[1] prygo> continue        # resumes execution (or exits if you’re in the REPL)
```

## Running Existing Code

For regular projects, you run the exact `go` command you already use but prefix it with `prygo`:

```bash
prygo test ./...
prygo run ./cmd/server
prygo build ./...
```

prygo scans the packages you requested, rewrites any files containing `pry.Pry()` to inject scope information, runs the underlying `go` command, then restores the original files. If your command fails (panic, crash, etc.), you can manually clean up with:

```bash
prygo revert
```

The `revert` subcommand walks the directories you targeted and restores files from the `.filename.gopry` backups.

## Example Snippets

### Inspecting State During a Test

```go
func TestCompute(t *testing.T) {
    got := compute(42)
    pry.Pry() // explore `got` + helper functions right before assertions
    if got != want {
        t.Fatalf("unexpected result")
    }
}
```

Run:

```bash
prygo test ./pkg/compute -run TestCompute
```

Evaluate expressions, poke at dependencies, then `continue` to finish the test.

### REPL Scratchpad

Use the `-generate` flag to produce a file you can inspect or commit:

```bash
prygo -i fmt,math -e 'fmt.Println(math.Sqrt(2))' -generate /tmp/pry-scratch.go
```

The file will already be instrumented with pry’s scope injection, so running `go run /tmp/pry-scratch.go` drops you straight into the prompt.

### Debugging CLI Flags

Pry works anywhere inside your code. For example:

```go
func configure() Config {
    cfg := load()
    pry.Pry()
    return cfg
}
```

Run:

```bash
prygo run ./cmd/tool --flag value
```

You can tweak fields interactively (`cfg.Timeout = time.Second`), call helper functions, then `continue` to see how the change behaves.

## Tips

- Use `continue` (or `exit`/`Ctrl+D`) in the prompt to resume program execution.
- History is persisted; press ↑/↓ to cycle previous expressions.
- `pry.Pry()` is inert if you run your program directly with `go`; the function is a no-op unless the injector replaces it.
- The generator automatically re-adds the original files after each command. If something fails mid-run, use `prygo revert` to restore.

With these patterns you can iteratively explore code, debug logic, or experiment with APIs without leaving Go’s toolchain. Happy pry’ing!
