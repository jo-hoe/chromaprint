package chromaprint

import (
	"os"
	"path/filepath"
	"runtime"
)

var (
	fpcalcFileName   = "fpcalc"
	windowsExtension = ".exe"
)

func NewBuilder() *builder {
	return &builder{
		filePath:             "",
		sampleRateInHz:       -1,
		channels:             -1,
		maxFingerPrintLength: -1,
		chunkSizeInSeconds:   -1,
		overlap:              false,
		algorithm:            -1,
	}
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

func (b *builder) WithPathToChromaprint(filePath string) *builder {
	b.filePath = filePath
	return b
}

func (b *builder) WithSampleRate(rateInHz int) *builder {
	b.sampleRateInHz = rateInHz
	return b
}

func (b *builder) WithChannels(numberOfChannels int) *builder {
	b.channels = numberOfChannels
	return b
}

func (b *builder) WithMaxFingerPrintLength(length int) *builder {
	b.maxFingerPrintLength = length
	return b
}

func (b *builder) WithChunksSize(chunkSizeInSeconds int) *builder {
	b.chunkSizeInSeconds = chunkSizeInSeconds
	return b
}

func (b *builder) WithAlgorithm(algorithm int) *builder {
	b.algorithm = algorithm
	return b
}

func (b *builder) WithOverlap(overlap bool) *builder {
	b.overlap = overlap
	return b
}

func (b *builder) Build() (*Chromaprint, error) {
	if b.filePath == "" {
		// on mac and linux file default name is "fpcalc"
		// and on windows it is"fpcalc.exe"
		chromaprintFilePath := fpcalcFileName
		if runtime.GOOS == "windows" {
			chromaprintFilePath = chromaprintFilePath + windowsExtension
		}

		folder, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		b.filePath = filepath.Join(folder, chromaprintFilePath)
	}

	return &Chromaprint{
		options: *b,
	}, nil
}
