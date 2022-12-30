package chromaprint

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Chromaprint struct {
	options builder
}

type Fingerprint struct {
	Timestamp   float64 `json:"timestamp"`
	Duration    float64 `json:"duration"`
        // TODO: convert into uint32
	Fingerprint []int   `json:"fingerprint"`
}

func (c *Chromaprint) CreateFingerprints(filepathToAudioFile string) ([]Fingerprint, error) {
	result := make([]Fingerprint, 0)

	if _, err := os.Stat(filepathToAudioFile); errors.Is(err, os.ErrNotExist) {
		return result, os.ErrNotExist
	}

	parameters := []string{filepathToAudioFile, "-json", "-raw"}
	parameters = append(parameters, c.getArgs()...)
	out, err := exec.Command(c.options.filePath, parameters...).Output()
	if err != nil {
		return result, err
	}

	// even if json format is chosen result value is not json
	// but a set of loose json elements
	// convert output into real json array of json element
	jsonString := strings.TrimSpace(string(out))
	// add comma separation between individual elements
	jsonString = strings.Replace(jsonString, "}", "},", -1)
	// cut last unneeded comma
	jsonString = jsonString[0 : len(jsonString)-1]
	// add array braces to whole set
	jsonString = fmt.Sprintf("[%s]", jsonString)

	err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return result, fmt.Errorf("invalid JSON output from fpcalc: %+v", err)
	}

	return result, nil
}

func (c *Chromaprint) getArgs() []string {
	args := make([]string, 0)

	addInt(&args, "rate", c.options.sampleRateInHz)
	addInt(&args, "channels", c.options.channels)
	addInt(&args, "length", c.options.maxFingerPrintLength)
	addInt(&args, "chunk", c.options.chunkSizeInSeconds)
	addInt(&args, "algorithm", c.options.algorithm)
	if c.options.overlap {
		args = append(args, "-overlap")
	}

	return args
}

func addInt(args *[]string, argName string, value int) {
	if value != -1 {
		*args = append(*args, fmt.Sprintf("-%s", argName), fmt.Sprint(value))
	}
}

func (c *Chromaprint) GetVersion() (string, error) {
	result, err := exec.Command(c.options.filePath, "-version").Output()
	return strings.TrimSpace(string(result)), err
}
