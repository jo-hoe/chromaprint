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

// Create a chromaprint builder
func NewBuilder() *builder {
	return &builder{
		filePath:             "",
		sampleRateInHz:       -1,
		channels:             -1,
		maxFingerPrintLength: 120,
		chunkSizeInSeconds:   -1,
		overlap:              false,
		algorithm:            2,
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

// Use this to specify the path to the Chromaprint CLI file.
func (b *builder) WithPathToChromaprint(filePath string) *builder {
	b.filePath = filePath
	return b
}

// Set the sample rate of the input audio
func (b *builder) WithSampleRate(rateInHz int) *builder {
	b.sampleRateInHz = rateInHz
	return b
}

// Set the number of channels in the input audio
func (b *builder) WithChannels(numberOfChannels int) *builder {
	b.channels = numberOfChannels
	return b
}

// Restrict the duration of the process input audio (default is 120)
func (b *builder) WithMaxFingerPrintLength(length int) *builder {
	b.maxFingerPrintLength = length
	return b
}

// Split the input audio into chunks of this duration
func (b *builder) WithChunksSize(chunkSizeInSeconds int) *builder {
	b.chunkSizeInSeconds = chunkSizeInSeconds
	return b
}

// Set the algorithm method (default 2)
func (b *builder) WithAlgorithm(algorithm int) *builder {
	b.algorithm = algorithm
	return b
}

// Overlap the chunks slightly to make sure audio on the edges is fingerprinted
func (b *builder) WithOverlap(overlap bool) *builder {
	b.overlap = overlap
	return b
}

// Builds the Chromaprint object
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
