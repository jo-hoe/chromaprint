package main

import (
	"fmt"

	"github.com/jo-hoe/chromaprint"
)

func main() {
	// assuming fpcalc.exe (aka chromaprint)
	// is in the same directory as this executable
	chromapint, err := chromaprint.NewBuilder().Build()
	if err != nil {
		fmt.Print(err)
		return
	}
	fingerprints, err := chromapint.CreateFingerprints("my.mp3")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%+v", fingerprints)
}
