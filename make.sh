#!/usr/bin/env

cd `dirname $0`
DIST_DIR_NAME="dist"
OS_LIST=("windows" "darwin" "linux")
ARCH_LIST=("386" "amd64")

set -eu

for OS in ${OS_LIST[@]}
do
    GOOS=${OS}
    for ARCH in ${ARCH_LIST[@]}
    do
        GOARCH=${ARCH}
        GODIST=${DIST_DIR_NAME}/${GOOS}/${GOARCH}
        mkdir -p ${GODIST}
        echo "compile "${GOOS}"("GOARCH")"
        go build -o gdbt main.go
        mv gdbt ${GODIST}/
    done
done