package chromaprint

import (
	"reflect"
	"runtime"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	filename := "fpcalc"
	if runtime.GOOS == "windows" {
		filename = filename + ".exe"
	}

	tests := []struct {
		name    string
		want    *builder
		wantErr bool
	}{
		{
			name: "init",
			want: &builder{
				filePath:             filename,
				sampleRateInHz:       -1,
				channels:             -1,
				maxFingerPrintLength: -1,
				chunkSizeInSeconds:   -1,
				algorithm:            -1,
				overlap:              false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBuilder()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBuilder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builder_WithFilePath(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		b    *builder
		args args
		want *builder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.WithFilePath(tt.args.filePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("builder.WithFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builder_WithSampleRate(t *testing.T) {
	type args struct {
		rateInHz int
	}
	tests := []struct {
		name string
		b    *builder
		args args
		want *builder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.WithSampleRate(tt.args.rateInHz); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("builder.WithSampleRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builder_WithChannels(t *testing.T) {
	type args struct {
		numberOfChannels int
	}
	tests := []struct {
		name string
		b    *builder
		args args
		want *builder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.WithChannels(tt.args.numberOfChannels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("builder.WithChannels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builder_WithMaxFingerPrintLength(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		b    *builder
		args args
		want *builder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.WithMaxFingerPrintLength(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("builder.WithMaxFingerPrintLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builder_WithChunksSize(t *testing.T) {
	type args struct {
		chunkSizeInSeconds int
	}
	tests := []struct {
		name string
		b    *builder
		args args
		want *builder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.WithChunksSize(tt.args.chunkSizeInSeconds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("builder.WithChunksSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builder_WithAlgorithm(t *testing.T) {
	type args struct {
		algorithm int
	}
	tests := []struct {
		name string
		b    *builder
		args args
		want *builder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.WithAlgorithm(tt.args.algorithm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("builder.WithAlgorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builder_WithOverlap(t *testing.T) {
	type args struct {
		overlap bool
	}
	tests := []struct {
		name string
		b    *builder
		args args
		want *builder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.WithOverlap(tt.args.overlap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("builder.WithOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_builder_Build(t *testing.T) {
	tests := []struct {
		name string
		b    *builder
		want *Chromaprint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("builder.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
