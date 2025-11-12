# Changelog

## v0.1.0 - 2025-11-12

### Highlights
- Reintroduces prygo as a maintained fork of go-pry with native support for Go 1.25 toolchains and modules.
- Mirrors the caller’s `go`/`toolchain` directives inside scratch workspaces so REPL sessions and `prygo run` behave like normal builds.
- Adds smarter import and module root detection so pry sessions automatically include interfaces, aliases, and local replaces.
- Ships an interactive default REPL plus the new `-generate` flag for writing pry-injected files without executing them.

### Tooling & Reliability
- Refreshed dependencies, cleaned the module graph, and added Makefile targets for linting, formatting, testing, and CI.
- Added GitHub Actions workflows (`ci.yml`, `unittest.yml`) that run lint, race-enabled tests, and build checks on Go 1.25.
- Documented recovery workflows (`prygo revert`) and improved debug logging to make injection failures easier to diagnose.

### Docs, Examples, and Playground
- Rewrote `README.md` and `examples.md` with installation, usage, and workflow guides plus animated demos.
- Added the `playground` and `example` programs used in screenshots/demos so users can try prygo without wiring their own project first.
- Captured the historical wiki content under `wiki_source/` and state-tracking fixtures for future tutorials.
