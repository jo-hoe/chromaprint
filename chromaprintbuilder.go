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
func NewBuilder() ChromaprintBuilder {
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
	// Builds the Chromaprint object
	Build() (*Chromaprint, error)
	// Use this to specify the path to the Chromaprint CLI file.
	WithPathToChromaprint(filePath string) ChromaprintBuilder
	// Set the sample rate of the input audio
	WithSampleRate(rateInHz int) ChromaprintBuilder
	// Set the number of channels in the input audio
	WithChannels(numberOfChannels int) ChromaprintBuilder
	// Restrict the duration of the process input audio (default is 120)
	WithMaxFingerPrintLength(length int) ChromaprintBuilder
	// Split the input audio into chunks of this duration
	WithChunksSize(chunkSizeInSeconds int) ChromaprintBuilder
	// Set the algorithm method (default 2)
	WithAlgorithm(algorithm int) ChromaprintBuilder
	// Overlap the chunks slightly to make sure audio on the edges is fingerprinted
	WithOverlap(overlap bool) ChromaprintBuilder
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

func (b *builder) WithPathToChromaprint(filePath string) ChromaprintBuilder {
	b.filePath = filePath
	return b
}

func (b *builder) WithSampleRate(rateInHz int) ChromaprintBuilder {
	b.sampleRateInHz = rateInHz
	return b
}

// Set the number of channels in the input audio
func (b *builder) WithChannels(numberOfChannels int) ChromaprintBuilder {
	b.channels = numberOfChannels
	return b
}

func (b *builder) WithMaxFingerPrintLength(length int) ChromaprintBuilder {
	b.maxFingerPrintLength = length
	return b
}

func (b *builder) WithChunksSize(chunkSizeInSeconds int) ChromaprintBuilder {
	b.chunkSizeInSeconds = chunkSizeInSeconds
	return b
}

func (b *builder) WithAlgorithm(algorithm int) ChromaprintBuilder {
	b.algorithm = algorithm
	return b
}

func (b *builder) WithOverlap(overlap bool) ChromaprintBuilder {
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
