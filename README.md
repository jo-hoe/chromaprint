# Chromaprint

[![GoDoc](https://godoc.org/github.com/jo-hoe/chromaprint?status.svg)](https://godoc.org/github.com/jo-hoe/chromaprint)
[![Test Status](https://github.com/jo-hoe/chromaprint/workflows/test/badge.svg)](https://github.com/jo-hoe/chromaprint/actions?workflow=test)
[![Coverage Status](https://coveralls.io/repos/github/jo-hoe/chromaprint/badge.svg?branch=main)](https://coveralls.io/github/jo-hoe/chromaprint?branch=main)
[![Lint Status](https://github.com/jo-hoe/chromaprint/workflows/lint/badge.svg)](https://github.com/jo-hoe/chromaprint/actions?workflow=lint)
[![CodeQL Status](https://github.com/jo-hoe/chromaprint/workflows/CodeQL/badge.svg)](https://github.com/jo-hoe/chromaprint/actions?workflow=CodeQL)
[![Go Report Card](https://goreportcard.com/badge/github.com/jo-hoe/chromaprint)](https://goreportcard.com/report/github.com/jo-hoe/chromaprint)

OS independent wrapper for [chromaprint](https://github.com/acoustid/chromaprint).
There exists already a [chromaprint API](https://github.com/go-fingerprint/gochroma).
However, it uses cgo and does not run on windows out of the box.
This wrapper removes the cgo dependency and allows to run chromaprint by using the binary directly.

## Example

```go
package main

import (
 "fmt"

 "github.com/jo-hoe/chromaprint"
)

func main() {
 // assuming fpcalc.exe (aka chromaprint)
 // is in the same directory as this executable
 chromaprinter, err := chromaprint.NewBuilder().Build()
 if err != nil {
  fmt.Print(err)
  return
 }
 fingerprints, err := chromaprinter.CreateFingerprints("my.mp3")
 if err != nil {
  fmt.Print(err)
  return
 }
 fmt.Printf("%+v", fingerprints)
}

```

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
