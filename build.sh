#!/bin/bash

set -e

echo "[ BUILD RELEASE ]"
BIN_DIR=$(pwd)/bin/
SUFFIX=""

BuildVersion=V1.0.0_$(git rev-parse --short HEAD | sed ':a;N;$!ba;s/^\n*//;s/\n*$//')
BuildTime=$(go run time.go)
BuildGoVersion=$(go version | sed ':a;N;$!ba;s/^\n*//;s/\n*$//')

# -ldflag 参数
GOLDFLAGS="-X 'main.BuildVersion=$BuildVersion'"
GOLDFLAGS+=" -X 'main.BuildTime=$BuildTime'"
GOLDFLAGS+=" -X 'main.BuildGoVersion=$BuildGoVersion'"

rm -rf "$BIN_DIR"
mkdir -p "$BIN_DIR"

cp *.toml "$BIN_DIR"

dist() {
    echo "try build GOOS=$1 GOARCH=$2"
    export GOOS=$g
    export GOARCH=$a
    export CGO_ENABLED=0
    if [ "$1" == "windows" ]; then
      SUFFIX=".exe"
    else
      SUFFIX=""
    fi
    go build -v -trimpath -ldflags "$GOLDFLAGS" -o "$BIN_DIR/easy-web-$1-$2$SUFFIX" "./cmd/easy-web"
    unset GOOS
    unset GOARCH
    echo "build success GOOS=$1 GOARCH=$2"
}

if [ "$1" == "dist" ]; then
    echo "[ DIST ALL PLATFORM ]"
    for g in "windows" "linux" "darwin"; do
        for a in "amd64" "386" "arm64" "arm"; do
            dist "$g" "$a"
        done
    done
else
  # build the current platform
  export GOOS=$(go env get GOOS | sed ':a;N;$!ba;s/^\n*//;s/\n*$//')
  export GOARCH=$(go env get GOARCH | sed ':a;N;$!ba;s/^\n*//;s/\n*$//')
  echo "[ DIST CURRENT PLATFORM GOOS=$GOOS GOARCH=$GOARCH]"
  if [ "$GOOS" == "windows" ]; then
    SUFFIX=".exe"
  fi
  go build -v -trimpath -ldflags "$GOLDFLAGS" -o "$BIN_DIR/easy-web-$GOOS-$GOARCH$SUFFIX" "./cmd/easy-web"
  echo "build success GOOS=$GOOS GOARCH=$GOARCH"
fi
