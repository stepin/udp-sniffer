.PHONY: all test build linux windows mac clean setup pre-release release help

PROJECTNAME=$(shell basename "$(PWD)")
BUILD_DATE=$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
GIT_COMMIT=$(shell git rev-parse HEAD)
PACKAGE_VERSION=local
export GO111MODULE := on

all: test build

## test: Run code reformat, tests, and lints.
test:
	go fmt
	go vet
	go test
	golangci-lint run --enable-all
	find . -name '*.go' | xargs grep -inHwE '(FIXME|TODO|HACK|XXX|BUG)' || true

## build: Build dev version for current OS.
build: deps
	go install -race -ldflags " -X main.version=${PACKAGE_VERSION} -X main.date=${BUILD_DATE} -X main.commit=${GIT_COMMIT}"

## linux: Build Linux test version.
linux: deps
	GOOS=linux go build -o ${PROJECTNAME} -ldflags " -X main.version=${PACKAGE_VERSION} -X main.date=${BUILD_DATE} -X main.commit=${GIT_COMMIT}"

## windows: Build Windows test version.
windows: deps
	GOOS=windows go build -o ${PROJECTNAME}.exe -ldflags " -X main.version=${PACKAGE_VERSION} -X main.date=${BUILD_DATE} -X main.commit=${GIT_COMMIT}"

## mac: Build macOS test version.
mac: deps
	GOOS=darwin go build -o ${PROJECTNAME} -ldflags ' -X main.version=${PACKAGE_VERSION} -X main.date=${BUILD_DATE} -X main.commit=${GIT_COMMIT}'

## clean: Clean folder.
clean:
	rm -rf ${PROJECTNAME} ${PROJECTNAME}.exe distr/ || true

## deps: Load dependent Go packages.
deps:
	go mod download

## setup: Download extra packages for development and build.
setup:
	#ide
	go get -v -u golang.org/x/tools/cmd/gorename
	go get -v -u github.com/rogpeppe/godef
	go get -v -u github.com/jstemmer/gotags
	go get -v -u github.com/nsf/gocode
	go get -v -u golang.org/x/tools/cmd/guru

	#use brew or other tool
	goreleaser -v
	golangci-lint -v

## pre-prelease: Test release procedure.
pre-release:
	goreleaser release --snapshot --skip-publish --rm-dist

## release: Package and upload to github releases and hub.docker.com. Don't forget to create branch with version name.
release:
	goreleaser release --rm-dist

## help: This help.
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo