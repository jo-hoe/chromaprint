package main

import (
	"fmt"

	"github.com/jo-hoe/chromaprint"
)

func main() {
	 // assuming fpcalc executable (aka chromaprint) is in the
	 // avaiable in $PATH. Alternativly use 
	 // NewBuilder().WithPathToChromaprint(/path/to/fpcalc)
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
