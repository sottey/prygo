# Troubleshooting prygo

This page collects solutions for common issues when running prygo locally or in CI.

## “updates to go.mod needed; run go mod tidy”

`go list` (called by `packages.Load`) detected stale module metadata. Run:

```bash
GOTOOLCHAIN=local go mod tidy
```

Commit the resulting `go.mod`/`go.sum`. Ensure your local Go version matches the `go` directive (e.g., Go 1.25 if `go 1.25` in `go.mod`), or the tool may downgrade and rewrite the directive.

## “invalid go version '1.24.0': must match format 1.23”

The `go` directive must be `major.minor` (e.g., `1.25`). Remove the patch component, then re-run `go mod tidy` using the desired toolchain (set `GOTOOLCHAIN=local` to prevent auto-downgrades).

## Missing History File in CI

If tests log `Failed to load the history ... no such file or directory`, they’re looking for `~/.prygo_history` on a clean runner. prygo now treats missing history as empty, but if you see this locally, create the file or ignore the warning—it’s non-fatal.

## prygo Reverting Files Too Early

`prygo run` automatically restores `.gopry` backups after the command finishes. If you want to keep the instrumented files (for debugging), pass `-r=false`:

```bash
prygo -r=false run ./cmd/app
```

Remember to run `prygo revert` afterwards to clean up.

## Toolchain Mismatch (Go 1.24 vs. 1.25)

prygo mirrors the module’s `go`/`toolchain` directives in temporary workspaces. If prygo can’t find the correct module root, set:

```bash
export PRYGO_MODULE_ROOT=/absolute/path/to/your/module
```

This ensures temp `go.mod` files reference the right toolchain. When editing `go.mod`, set `GOTOOLCHAIN=local` so the go tool uses your installed compiler instead of auto-downloading an older toolchain.

## “operation not permitted” when running `go mod tidy`

In sandboxed environments (like codespaces or Codex CLI), writes to the default Go build cache can be blocked. Set `GOCACHE` and `GOMODCACHE` to directories you control, e.g.:

```bash
export GOCACHE=$PWD/.cache/go-build
export GOMODCACHE=$PWD/.cache/pkg/mod
mkdir -p "$GOCACHE" "$GOMODCACHE"
```

Then rerun `go mod tidy`.

## Debug Logging (-d)

If injection fails silently, rerun with debug logs:

```bash
prygo -d run ./cmd/app
```

The logs will show which files were rewritten, how many `pry.Pry()` statements were found, and any errors encountered.

Need help with an issue not covered here? Open an issue on GitHub or update this page with the fix you discover.
