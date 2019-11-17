# slate

[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/caigwatkin/slate/master/LICENSE)
[![CircleCI](https://circleci.com/gh/caigwatkin/slate/tree/master.svg?style=svg&circle-token=3e3e4785268df5c0bcfafacdca9c4717b5d36317)](https://circleci.com/gh/caigwatkin/slate/tree/master)
[![codecov](https://codecov.io/gh/caigwatkin/slate/branch/master/graph/badge.svg?token=OL7VqVYJLU)](https://codecov.io/gh/caigwatkin/slate)
[![Go Report Card](https://goreportcard.com/badge/github.com/caigwatkin/slate)](https://goreportcard.com/report/github.com/caigwatkin/slate)

An API server written in Go. A clean slate, if you will.

## Usage

```bash
go get -u github.com/caigwatkin/slate
slate -h
```

## CI/CD

Using [CircelCI](https://circleci.com) for builds of commits and pull requests.

All changes are made to branches of `master`. The branch must be up to date with `master` and all commits must be signed with a [GPG key](https://gnupg.org).

The following status checks must pass before merging into master:

- [CircelCI](https://circleci.com) build passes
- [Codecov](https://codecov.io) meets minimum coverage requirements
