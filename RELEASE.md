# Release guide

These are the steps I followed to get `v0.1.0` ready. Re-run the checklist for every future tag so GitHub releases stay reproducible.

## 1. Prep work
1. Update `CHANGELOG.md` with a new section and date.
2. Run the fast feedback loop locally (or rely on Actions):
   ```bash
   export GOMODCACHE=$(pwd)/.gocache/mod
   make ci
   ```
   `GOMODCACHE` keeps downloads inside the repo so the sandbox does not block toolchain installs.
3. Spot-check `README.md` / `examples.md` for version-specific language.

## 2. Build release artifacts
Create a clean `dist/` folder and compile the CLI for the platforms you want to attach to the release.
```bash
rm -rf dist && mkdir dist
for combo in \
  "darwin arm64" \
  "darwin amd64" \
  "linux arm64" \
  "linux amd64" \
  "windows amd64"
do
  GOOS=${combo% *}
  GOARCH=${combo#* }
  SUFFIX=""
  [[ "$GOOS" == "windows" ]] && SUFFIX=".exe"
  GOMODCACHE=$(pwd)/.gocache/mod \
  GOOS=$GOOS GOARCH=$GOARCH \
    go build -trimpath -ldflags "-s -w" \
    -o dist/prygo-${GOOS}-${GOARCH}${SUFFIX} .
done
```
Compress each binary (zip for Windows, tar.gz elsewhere) before uploading.
```bash
cd dist
tar -czf prygo-darwin-arm64.tar.gz prygo-darwin-arm64
tar -czf prygo-darwin-amd64.tar.gz prygo-darwin-amd64
tar -czf prygo-linux-arm64.tar.gz prygo-linux-arm64
tar -czf prygo-linux-amd64.tar.gz prygo-linux-amd64
zip -r prygo-windows-amd64.zip prygo-windows-amd64.exe
cd -
```

## 3. Tag the release
```bash
git tag -a v0.1.0 -m "v0.1.0"
git push origin v0.1.0
```

## 4. Publish on GitHub
1. Open **Releases → Draft a new release**.
2. Choose the `v0.1.0` tag (or create it), set the title to `prygo v0.1.0`.
3. Paste the `CHANGELOG` entry as the notes.
4. Upload the archives from `dist/`.
5. Click **Publish release**.

## 5. Post-release
- Update `README` installation snippets if you want to highlight the new tag (`go install github.com/sottey/prygo@v0.1.0`).
- Announce the release (blog, X, Mastodon, Slack, etc.).
