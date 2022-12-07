package chromaprint

import (
	"context"
	"errors"
	"os"
)

type ChromaprintCtx struct {
	filePath             string
	sampleRateInHz       int
	channels             int
	maxFingerPrintLength int
	chunkSizeInSeconds   int
	overlap              bool
	algorithm            int
}

// initalizes chromaprint context with default parameters
//
// filePath is the path to the chromaprint executable
func NewChromaprintCtx(filePath string) (*ChromaprintCtx, error) {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return nil, os.ErrNotExist
	}

	context.Background()

	return &ChromaprintCtx{
		filePath:             filePath,
		sampleRateInHz:       -1,
		channels:             -1,
		maxFingerPrintLength: -1,
		chunkSizeInSeconds:   -1,
		overlap:              false,
		algorithm:            -1,
	}, nil
}

// should follow
// https://mauricio.github.io/2022/02/07/gof-patterns-in-golang.html
// or https://golang.cafe/blog/golang-functional-options-pattern.html
// option pattern for immutables
// builder for changing option
