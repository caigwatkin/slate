#!/bin/bash -eu

# Run:
# $ ./start.sh [cmd]

SCRIPT_DIR="$( cd "$(dirname "$0")" ; pwd -P )"
SCRIPT_DIR_NAME=${SCRIPT_DIR##*/}
SCRIPT_NAME=`basename $0`

API=${1}
API_MAIN=./cmd/${1}

export GO111MODULE=on

echo "${SCRIPT_NAME} -> START at `date '+%Y-%m-%d %H:%M:%S'`..."

echo "Formatting..."
go fmt ${API_MAIN}/...
echo

echo "Vetting..."
go vet -mod=vendor ${API_MAIN}/...
echo

echo "Linting..."
revive -formatter stylish -config ${SCRIPT_DIR}/../../config/revive.toml -exclude ${SCRIPT_DIR}/../../vendor/... ${SCRIPT_DIR}/...
echo

echo "Testing..."
go test -mod=vendor ${API_MAIN}/...
echo

echo "Building..."
go build -mod=vendor -o bin/${API} ${API_MAIN}
echo

echo "Running..."
./bin/${API}
