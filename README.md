# slate

[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/caigwatkin/slate/master/LICENSE)
[![Build Status](https://travis-ci.org/caigwatkin/slate.svg?branch=master)](https://travis-ci.org/caigwatkin/slate)
[![codecov](https://codecov.io/gh/caigwatkin/slate/branch/master/graph/badge.svg)](https://codecov.io/gh/caigwatkin/slate)
[![GolangCI](https://golangci.com/badges/github.com/caigwatkin/slate.svg)](https://golangci.com)
[![Go Report Card](https://goreportcard.com/badge/github.com/caigwatkin/slate)](https://goreportcard.com/report/github.com/caigwatkin/slate)

An API server written in Go. A clean slate, if you will.

## Usage

Start the server by running the following from repo root:

```bash
./scripts/go/slate.sh
```

## Project structure

Uses project structure/layout as recommended by [golang standards guide](https://github.com/golang-standards/project-layout)

## CI/CD

Using [Travis CI](https://travis-ci.org) for builds of commits and pull requests.

All changes are made to branches of `master`. The branch must be up to date with `master` and all commits must be signed with a [GPG key](https://gnupg.org).

The following status checks must pass before merging into master:

- [Travis CI](https://travis-ci.org) build passes
- [Codecov](https://codecov.io) meets minimum coverage requirements
- [GolangCI](https://golangci.com) finds no issues

## Dependency management

Using [Go 1.11 Modules](https://github.com/golang/go/wiki/Modules)
