#!/usr/bin/env

cd `dirname $0`
SRC_PATH=`pwd`
DIST_DIR_NAME="dist"
OS_LIST=("windows" "darwin" "linux")
ARCH_LIST=("386" "amd64")

set -eu

rm -rf ${DIST_DIR_NAME}
for OS in ${OS_LIST[@]}
do
    GOOS=${OS}
    for ARCH in ${ARCH_LIST[@]}
    do
        GOARCH=${ARCH}
        GODIST=${DIST_DIR_NAME}/${GOOS}/${GOARCH}
        mkdir -p ${GODIST}
        echo "compile "${GOOS}"("${GOARCH}")"
        cd ${GODIST}
        env GOOS=${GOOS} GOARCH=${GOARCH} go build -v ${SRC_PATH}/gdbt.go
        cd ${SRC_PATH}
    done
done