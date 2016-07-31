#!/bin/bash
set -e
set -x
cd "$(dirname "$0")"

export GO15VENDOREXPERIMENT=1
go fmt
go vet
go test $(glide novendor)
golint ./...
gocyclo -over 25 .
interfacer ./...
find . -name '*.go' | xargs grep -inHwE '(FIXME|TODO|HACK|XXX|BUG)' || true
errcheck -verbose ./...
