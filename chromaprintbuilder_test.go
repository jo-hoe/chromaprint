package chromaprint

import (
	"reflect"
	"runtime"
	"testing"
)

func Test_builder_Build(t *testing.T) {
	filename := "fpcalc"
	if runtime.GOOS == "windows" {
		filename = filename + ".exe"
	}

	tests := []struct {
		name string
		b    *builder
		want *Chromaprint
	}{
		{
			name: "default",
			b:    newBuilder(t),
			want: &Chromaprint{
				options: builder{
					filePath:             filename,
					sampleRateInHz:       -1,
					channels:             -1,
					maxFingerPrintLength: -1,
					chunkSizeInSeconds:   -1,
					algorithm:            -1,
					overlap:              false,
				},
			},
		}, {
			name: "full bild",
			b:    newBuilder(t).WithFilePath("test").WithSampleRate(44100).WithChannels(2).WithMaxFingerPrintLength(120).WithChunksSize(5).WithAlgorithm(1).WithOverlap(true),
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("builder.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newBuilder(t *testing.T) *builder {
	result, err := NewBuilder()
	if err != nil {
		t.Error(err)
	}
	return result
}
