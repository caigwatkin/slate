#!/bin/bash -eu

# Run from repo root:
# $ ./scripts/go/start.sh [cmd]

SCRIPT_DIR="$( cd "$(dirname "$0")" ; pwd -P )"
SCRIPT_DIR_NAME=${SCRIPT_DIR##*/}
SCRIPT_NAME=`basename $0`

export GO111MODULE=on

echo "${SCRIPT_NAME} -> START at `date '+%Y-%m-%d %H:%M:%S'`..."

cd ${SCRIPT_DIR}/../..

echo "Formatting..."
go fmt ./...
echo

echo "Vetting..."
go vet -mod=vendor ./...
echo

echo "Linting..."
revive -formatter=stylish -config=./configs/revive.toml -exclude=./vendor/... ./...
echo

echo "Testing..."
go test -mod=vendor ./...
echo

echo "Running..."
go run -mod=vendor ./cmd/service
