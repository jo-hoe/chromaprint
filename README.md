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

```golang
package main

import (
 "fmt"

 "github.com/jo-hoe/chromaprint"
)

func main() {
 // assuming fpcalc executable (aka chromaprint) is in the
 // avaiable in $PATH. Alternativly use 
 // NewBuilder().WithPathToChromaprint("/path/to/fpcalc")
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

## Chromaprint

Chromaprint provides a fingerprint for a given audio.
The aim is to provide a single identifier for a given audio input.
This identifier should be robust to background noice.

Here is how it roughly works:

- audio is resampled to mono at 11025Hz
- the result is cut into individual frames
- fast fourier transform is performed to extract frequencies within each chucks
- features are extracted from the these frequencies 
- these features are converted into an image
- a fingerprint is calculated from this image

### Chromaprint Algorithms

Chromaprint allows to set an algorithm for the generation of fingerprint.

```
// where n is an integer between 1 and 5
chromaprint.NewBuilder().WithAlgorithm(n)
```

Each of these provides a different fingerprint for a given audio input.
Each algorithm requires a minimum amount of audio data before a fingerprint can be calculated.
Below is a table showing how long an audio clip has to be before the different algorithms can provide a fingerprint.

|# Algorithm|Minimum seconds of audio|
|---|---|
|1|~ 3|
|2|~ 3|
|3|~ 7.5|
|4|~ 3.5|
|5|~ 2|

Note, that these values are only rough approximations and differ depending on the audio input.

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
