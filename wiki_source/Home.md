# prygo Wiki

Welcome to the prygo wiki! This is the jumping-off point for learning what prygo does, how to install it, and how to get the most out of the interactive REPL.

## What is prygo?

prygo is a drop-in wrapper around the Go toolchain that injects `pry.Pry()` breakpoints into your code. When execution reaches a pry statement you get a live REPL with every variable and imported package already in scope. It’s heavily inspired by Ruby’s Pry, but tailored for Go’s compiled workflow.

- Supports Go 1.25+ with module-aware temp workspaces.
- Mirrors your project’s `go`/`toolchain` directives so pry sessions use the same compiler as your build.
- Works with `go run`, `go test`, `go build`, or a standalone REPL (`prygo` with no args).

## Quick Links

- **Getting Started** – [Basics and installation](./Getting-Started.md)
- **Examples & Workflows** – [Hands-on scenarios, REPL tricks, and advanced flags](./Examples.md)
- **Troubleshooting** – [Common issues, `prygo revert`, temp module fixes, and environment tips](./Troubleshooting.md)
- **Contributing** – [Development workflow, CI, linting, and code style](./Contributing.md)

## Frequently Used Commands

```bash
# Install
go install github.com/sottey/prygo@latest

# Run your app with pry breakpoints
prygo run ./cmd/app

# Drop into a scratch REPL
prygo -i fmt,net/http

# Revert injected files if a run crashes
prygo revert ./path/to/package
```

For a deeper dive, jump to the [Examples](./Examples.md) page, which mirrors the `examples.md` file in the repo and expands on practical debugging workflows.
