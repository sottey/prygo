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
