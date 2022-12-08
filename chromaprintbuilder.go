package chromaprint

import (
	"errors"
	"os"
)

func NewBuilder(chromaprintFilePath string) (*builder, error) {
	if _, err := os.Stat(chromaprintFilePath); errors.Is(err, os.ErrNotExist) {
		return nil, os.ErrNotExist
	}

	return &builder{
		filePath:             chromaprintFilePath,
		sampleRateInHz:       -1,
		channels:             -1,
		maxFingerPrintLength: -1,
		chunkSizeInSeconds:   -1,
		overlap:              false,
		algorithm:            -1,
	}, nil
}

// ChromaprintBuilder defines the fields we want to set on this builder, you could add/remove
// fields here.
type ChromaprintBuilder interface {
	Build() (*Chromaprint, error)
}

type builder struct {
	filePath             string
	sampleRateInHz       int
	channels             int
	maxFingerPrintLength int
	chunkSizeInSeconds   int
	overlap              bool
	algorithm            int
}

func (b *builder) Build() *Chromaprint {
	return &Chromaprint{
		options: *b,
	}
}