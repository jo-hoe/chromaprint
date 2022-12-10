# Chromaprint

[![GoDoc](https://godoc.org/github.com/jo-hoe/chromaprint?status.svg)](https://godoc.org/github.com/jo-hoe/chromaprint)
[![Test Status](https://github.com/jo-hoe/chromaprint/workflows/test/badge.svg)](https://github.com/jo-hoe/chromaprint/actions?workflow=test)
[![Coverage Status](https://coveralls.io/repos/github/jo-hoe/chromaprint/badge.svg?branch=main)](https://coveralls.io/github/jo-hoe/chromaprint?branch=main)
[![Lint Status](https://github.com/jo-hoe/chromaprint/workflows/lint/badge.svg)](https://github.com/jo-hoe/chromaprint/actions?workflow=lint)
[![CodeQL Status](https://github.com/jo-hoe/chromaprint/workflows/CodeQL/badge.svg)](https://github.com/jo-hoe/chromaprint/actions?workflow=CodeQL)
[![Go Report Card](https://goreportcard.com/badge/github.com/jo-hoe/chromaprint)](https://goreportcard.com/report/github.com/jo-hoe/chromaprint)

OS independent wrapper for [chromaprint](https://github.com/acoustid/chromaprint).
There exists already a [chromaprint API](https://github.com/go-fingerprint/gochroma).
However, it is not interperable and does not run on windows.

## Dependency

It is required have a [chromaprint](https://acoustid.org/chromaprint) CLI.

## License

While this repository uses MIT the chromaprint license may differ (i.e. LGPL2.1).

## Linting

Project used `golangci-lint` for linting.

### Installation

<https://golangci-lint.run/usage/install/>

### Execution

Run the linting locally by executing

```cli
golangci-lint run ./...
```

in the working directory
