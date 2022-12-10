package main

import (
	"fmt"

	"github.com/jo-hoe/chromaprint"
)

func main() {
	// assuming both fpcalc.exe (aka chromaprint)
	// is in the same directory as this executable
	builder, err := chromaprint.NewBuilder()
	if err != nil {
		fmt.Print(err)
		return
	}
	chromapint := builder.Build()
	fingerprints, err := chromapint.CreateFingerprints("my.mp3")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%+v", fingerprints)
}
