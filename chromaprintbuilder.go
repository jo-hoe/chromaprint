package chromaprint

import "runtime"

var (
	fpcalcFileName   = "fpcalc"
	windowsExtension = ".exe"
)

func NewBuilder() (*builder, error) {
	// on mac and linux file default name is "fpcalc"
	// and on windows it is"fpcalc.exe"
	chromaprintFilePath := fpcalcFileName
	if runtime.GOOS == "windows" {
		chromaprintFilePath = chromaprintFilePath + windowsExtension
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
