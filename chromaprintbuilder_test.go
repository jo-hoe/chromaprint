package chromaprint

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func Test_builder_Build(t *testing.T) {
	filename := "fpcalc"
	if runtime.GOOS == "windows" {
		filename = filename + ".exe"
	}

	path, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	filePath := filepath.Join(path, filename)

	tests := []struct {
		name    string
		b       *builder
		want    *Chromaprint
		wantErr bool
	}{
		{
			name: "default build",
			b:    NewBuilder(),
			want: &Chromaprint{
				options: builder{
					filePath:             filePath,
					sampleRateInHz:       -1,
					channels:             -1,
					maxFingerPrintLength: -1,
					chunkSizeInSeconds:   -1,
					algorithm:            -1,
					overlap:              false,
				},
			},
			wantErr: false,
		}, {
			name: "full build",
			b:    NewBuilder().WithPathToChromaprint("test").WithSampleRate(44100).WithChannels(2).WithMaxFingerPrintLength(120).WithChunksSize(5).WithAlgorithm(1).WithOverlap(true),
			want: &Chromaprint{
				options: builder{
					filePath:             "test",
					sampleRateInHz:       44100,
					channels:             2,
					maxFingerPrintLength: 120,
					chunkSizeInSeconds:   5,
					algorithm:            1,
					overlap:              true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.Build()
			if (err != nil) != tt.wantErr {
				t.Errorf("Chromaprint.CreateFingerprints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("builder.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
