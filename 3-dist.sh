#!/bin/bash
set -e
cd "$(dirname "$0")"

export PACKAGE_NAME=udp-sniffer
export PACKAGE_VERSION=1.0
export BUILD_DATE=`date -u '+%Y-%m-%d_%I:%M:%S%p'`

export PKG_DIR="`pwd`/pkg"
export DIST_DIR="`pwd`/dist"

#wercker compatible variables
export WERCKER_GIT_COMMIT=`git rev-parse HEAD`
export WERCKER_GIT_OWNER=`git config github.username`
export WERCKER_GIT_REPOSITORY=`git rev-parse --show-toplevel|xargs basename`
export GITHUB_TOKEN=`git config github.token`


echo Clean build folders
rm -rf $PKG_DIR $DIST_DIR

echo Cross-compile
gox -osarch='darwin/amd64 linux/386 linux/amd64 windows/386 windows/amd64' -output "pkg/$PACKAGE_NAME-{{.OS}}_{{.Arch}}/{{.Dir}}" -ldflags "-X main.version=$PACKAGE_VERSION -X main.buildDate=$BUILD_DATE -X main.gitCommit=$WERCKER_GIT_COMMIT"

echo Copy text files to distr
for TARGET in $(find ${PKG_DIR} -mindepth 1 -maxdepth 1 -type d); do
    cp README.* LICENSE ${TARGET}
done

echo Zip folders
mkdir -p ${DIST_DIR}
for TARGET in $(find ${PKG_DIR} -mindepth 1 -maxdepth 1 -type d); do
    ARCHIVE_NAME=$(basename ${TARGET})
    pushd ${TARGET}
    zip -r ${DIST_DIR}/${ARCHIVE_NAME}.zip ./*
    popd
done
echo Generate shasum
pushd ${DIST_DIR}
shasum * > ./SHASUMS
popd

#echo Create rpm
#fpm -s dir -t rpm --name $PACKAGE_NAME --version $PACKAGE_VERSION --package dist/ --license MIT --rpm-os linux --prefix /opt/$PACKAGE_NAME-$PACKAGE_VERSION --url https://github.com/stepin/udp-sniffer pkg/$PACKAGE_NAME-linux_amd64/

echo Upload release to github
#ghr --token $GITHUB_TOKEN --username $WERCKER_GIT_OWNER --repository $WERCKER_GIT_REPOSITORY --replace `git describe --tags` dist/
