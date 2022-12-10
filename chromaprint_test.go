package chromaprint

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestChromaprint_CreateFingerprints(t *testing.T) {
	type args struct {
		filepathToAudioFile string
	}
	tests := []struct {
		name    string
		c       *Chromaprint
		args    args
		want    []Fingerprint
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CreateFingerprints(tt.args.filepathToAudioFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chromaprint.CreateFingerprints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chromaprint.CreateFingerprints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChromaprint_getArgs(t *testing.T) {
	tests := []struct {
		name string
		c    *Chromaprint
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.getArgs(); got != tt.want {
				t.Errorf("Chromaprint.getArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addInt(t *testing.T) {
	type args struct {
		builder *strings.Builder
		argName string
		value   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addInt(tt.args.builder, tt.args.argName, tt.args.value)
		})
	}
}

func TestChromaprint_GetVersion(t *testing.T) {
	tests := []struct {
		name    string
		c       *Chromaprint
		want    string
		wantErr bool
	}{
		{
			name:    "get version",
			c:       getDefaultChromaprint(t),
			want:    "fpcalc version 1.5.1 (FFmpeg Lavc58.134.100 Lavf58.76.100 SwR3.9.100)",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetVersion()
			if (err != nil) != tt.wantErr {
				t.Errorf("Chromaprint.GetVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Chromaprint.GetVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getDefaultChromaprint(t *testing.T) *Chromaprint {
	// get os dependend file name
	filename := "fpcalc"
	if runtime.GOOS == "windows" {
		filename = "win-" + filename + ".exe"
	} else if runtime.GOOS == "darwin" {
		filename = "macos-" + filename
	} else {
		filename = "linux-" + filename
	}

	// get test resource folder
	testFileFolder, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	filePath := filepath.Join(testFileFolder, "testresources", filename)

	result := Chromaprint{
		options: builder{
			filePath: filePath,
		},
	}
	return &result
}
