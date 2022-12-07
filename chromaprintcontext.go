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
// https://dev.to/gopher/getting-started-with-go-context-l7g
// or https://golang.cafe/blog/golang-functional-options-pattern.html
