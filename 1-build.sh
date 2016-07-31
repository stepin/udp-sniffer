#!/bin/bash
set -e
set -x
cd "$(dirname "$0")"

export BUILD_DATE=`date -u '+%Y-%m-%d_%I:%M:%S%p'`
export WERCKER_GIT_COMMIT=`git rev-parse HEAD`
export PACKAGE_VERSION=local

export GO15VENDOREXPERIMENT=1
go install -race -ldflags " -X main.version=$PACKAGE_VERSION -X main.buildDate=$BUILD_DATE -X main.gitCommit=$WERCKER_GIT_COMMIT"
