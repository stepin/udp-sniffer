#!/bin/bash
set -e
cd "$(dirname "$0")"

go fmt
go vet
go test $(glide novendor)
golint ./...
gocyclo -over 25 .
interfacer ./...
find . -name '*.go' | xargs grep -inHwE '(FIXME|TODO|HACK|XXX|BUG)'
