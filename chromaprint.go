package chromaprint

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Chromaprint struct {
	options builder
}

type Fingerprint struct {
	Duration    float64 `json:"duration"`
	Fingerprint []int32 `json:"fingerprint"`
}

func (c *Chromaprint) CreateFingerprints(filepathToAudioFile string) ([]Fingerprint, error) {
	result := make([]Fingerprint, 0)

	if _, err := os.Stat(filepathToAudioFile); errors.Is(err, os.ErrNotExist) {
		return nil, os.ErrNotExist
	}

	out, err := exec.Command(c.options.filePath, filepathToAudioFile, "-json", "-raw", c.getArgs()).Output()
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(out, &result)
	if err != nil {
		return result, fmt.Errorf("invalid JSON output from fpcalc: %w", err)
	}

	return result, nil
}

func (c *Chromaprint) getArgs() string {
	var stringBuilder strings.Builder

	addInt(&stringBuilder, "channels", c.options.channels)
	addInt(&stringBuilder, "algorithm", c.options.algorithm)
	addInt(&stringBuilder, "chunk", c.options.chunkSizeInSeconds)
	addInt(&stringBuilder, "length", c.options.maxFingerPrintLength)
	addInt(&stringBuilder, "rate", c.options.sampleRateInHz)
	if c.options.overlap {
		_, err := stringBuilder.WriteString("-overlap")
		log.Printf("%+v", err)
	}

	return stringBuilder.String()
}

func addInt(builder *strings.Builder, argName string, value int) {
	if value != -1 {
		_, err := builder.WriteString(fmt.Sprintf("-%s ", argName))
		log.Printf("%+v", err)
		_, err = builder.WriteString(fmt.Sprint(value))
		log.Printf("%+v", err)
	}
}

func (c *Chromaprint) GetVersion() (string, error) {
	result, err := exec.Command(c.options.filePath, "-version").Output()
	return strings.TrimSpace(string(result)), err
}
