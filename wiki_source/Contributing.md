# Contributing to prygo

Thanks for your interest in improving prygo! This guide outlines how to set up your environment, run quality checks, and submit changes.

## Development Environment

1. Install Go 1.25.x (or the version declared in `go.mod`).
2. Install `golangci-lint`:
   ```bash
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   ```
3. Clone the repo and ensure `make ci` succeeds:
   ```bash
   git clone https://github.com/sottey/prygo.git
   cd prygo
   make ci
   ```

Targets in the `Makefile`:

| Target | Description |
|--------|-------------|
| `make build` | `go build ./...` |
| `make test`  | `go test -count=1 ./...` |
| `make lint`  | `golangci-lint run ./...` |
| `make ci`    | Runs lint + test (used in GitHub Actions) |

## Coding Guidelines

- Follow standard Go formatting (`go fmt ./...`).
- Keep comments concise; add them only when code isn’t self-explanatory.
- Prefer small, focused commits; if you’re changing multiple areas, split them logically.
- Avoid introducing new dependencies unless necessary. If you must, explain why in the PR.

## Testing

- Run `make ci` before opening a pull request.
- For changes touching the REPL or file injection, add/adjust tests in `generate/` or `pry/`.
- Example directories under `example/` are excluded from CI builds via `//go:build ignore`. If you add a runnable example, ensure it compiles without extra setup.

## Documentation

- Update `README.md`, `examples.md`, and relevant wiki pages (under `wiki_source/`) when adding features or changing workflows.
- Mention new flags or behaviours in `flag.CommandLine.Usage` (in `main.go`).

## Pull Requests

1. Fork the repository and create a feature branch.
2. Make your changes, keeping commits tidy.
3. Run `make ci`.
4. Open a PR against `master`, describing the change and any testing performed.
5. Address review feedback promptly; we aim for an inclusive, collaborative process.

## Code of Conduct

Be respectful, constructive, and helpful. prygo’s community is small but friendly—help newcomers, document your discoveries, and treat others as teammates.

If you have ideas for large changes, open an issue first so we can discuss design tradeoffs before you invest time in implementation.
