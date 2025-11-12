# prygo Examples & Workflows

This page mirrors and expands on `examples.md` from the repository. Use it as a reference for common debugging scenarios.

## Simple Breakpoint in `main()`

```go
import "github.com/sottey/prygo/pry"

func main() {
    total := compute()
    pry.Pry()
    fmt.Println(total)
}
```

Run:

```bash
prygo run ./cmd/app
```

When the program hits `pry.Pry()`, evaluate expressions (e.g., `total`, `helpers.Inspect(total)`) or call functions before continuing.

## Debugging Tests

Insert a pry in a test to check intermediate state:

```go
func TestPipeline(t *testing.T) {
    result := buildPipeline()
    pry.Pry()
    if result.Err != nil {
        t.Fatal(result.Err)
    }
}
```

```bash
prygo test ./pkg/pipeline -run TestPipeline
```

You can mutate `result` or call helpers to reproduce race conditions before letting the test proceed.

## Standalone REPL Scratchpad

Generate a pry-injected file for ad-hoc experiments:

```bash
prygo -i fmt,math -e 'fmt.Println(math.Sqrt(2))' -generate /tmp/pry-scratch.go
```

Run it later with `go run /tmp/pry-scratch.go` and you’ll drop directly into the REPL.

## Inspecting CLI Flags

Add pry inside initialization logic:

```go
func configure() Config {
    cfg := loadFlags()
    pry.Pry()
    return cfg
}
```

```bash
prygo run ./cmd/tool --flag value
```

Check `cfg`, tweak settings (e.g., `cfg.Timeout = time.Second`) and see how the rest of the program reacts.

## Using `prygo revert`

If a pry session crashes before cleanup, run:

```bash
prygo revert ./pkg/foo
```

This restores any `.file.gopry` backups to their original names. It’s safe to run multiple times.

## Tips

- `continue` or `exit` leaves the REPL; `Ctrl+D` also exits.
- `history` is stored in `~/.prygo_history`; prygo handles missing files gracefully.
- Combine prygo with `go test ./...` to trace failing packages individually.
- Set `PRYGO_MODULE_ROOT` when working in nested directories so prygo can locate the right module metadata.

Have a workflow worth sharing? Add it to `examples.md` in the repo and mirror it here!
