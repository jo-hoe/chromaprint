package chromaprint

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
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

	out, err := exec.Command(c.options.filePath, filepathToAudioFile, "-json", "-raw").Output()
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(out, &result)
	if err != nil {
		return result, fmt.Errorf("invalid JSON output from fpcalc: %w", err)
	}

	return result, nil
}

func (c *Chromaprint) GetVersion() string {
	return ""
}
