package chromaprint

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	testMP3             = "edgar-allen-poe-the-telltale-heart-original.mp3"
	testMP3_Folder_Name = "mp3"
)

func TestChromaprint_CreateFingerprints(t *testing.T) {
	type args struct {
		filepathToAudioFile string
	}
	tests := []struct {
		name    string
		c       *Chromaprint
		args    args
		want    []FingerprintData
		wantErr bool
	}{
		{
			name: "no executable",
			c:    &Chromaprint{},
			args: args{
				filepathToAudioFile: filepath.Join(getTestFolderPath(t), testMP3),
			},
			want:    make([]FingerprintData, 0),
			wantErr: true,
		}, {
			name: "no mp3",
			c: &Chromaprint{
				options: builder{
					filePath: getExecutable(t),
				},
			},
			args: args{
				filepathToAudioFile: "none_existing.mp3",
			},
			want:    make([]FingerprintData, 0),
			wantErr: true,
		}, {
			name: "wrong config",
			c: &Chromaprint{
				options: builder{
					filePath:             getExecutable(t),
					maxFingerPrintLength: 0,
				},
			},
			args: args{
				filepathToAudioFile: filepath.Join(getTestFolderPath(t), testMP3),
			},
			want:    make([]FingerprintData, 0),
			wantErr: true,
		}, {
			name: "default positiv test",
			c: &Chromaprint{
				options: builder{
					filePath:             getExecutable(t),
					sampleRateInHz:       -1,
					channels:             -1,
					maxFingerPrintLength: 3,
					chunkSizeInSeconds:   -1,
					overlap:              false,
					algorithm:            -1,
				},
			},
			args: args{
				filepathToAudioFile: filepath.Join(getTestFolderPath(t), testMP3_Folder_Name, testMP3),
			},
			want: []FingerprintData{{
				DurationInSeconds: 1059.97,
				Fingerprint:       []uint32{1920772148, 1932307492, 1999416352},
			}},
			wantErr: false,
		}, {
			name: "chucked",
			c: &Chromaprint{
				options: builder{
					filePath:             getExecutable(t),
					sampleRateInHz:       -1,
					channels:             -1,
					maxFingerPrintLength: 6,
					chunkSizeInSeconds:   3,
					overlap:              false,
					algorithm:            -1,
				},
			},
			args: args{
				filepathToAudioFile: filepath.Join(getTestFolderPath(t), testMP3_Folder_Name, testMP3),
			},
			want: []FingerprintData{{
				TimestampInSeconds: 0,
				DurationInSeconds:  3,
				Fingerprint:        []uint32{1920772148, 1932307492, 1999416352},
			}, {
				TimestampInSeconds: 3,
				DurationInSeconds:  3,
				Fingerprint:        []uint32{1390452773, 1390455845, 1398844461},
			}},
			wantErr: false,
		},
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

func Test_addInt(t *testing.T) {
	actual := make([]string, 0)
	addInt(&actual, "test", 1)

	expected := []string{"-test", "1"}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected '%s' actual '%s'", expected, actual)
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

func getTestFolderPath(t *testing.T) string {
	// get test resource folder
	testFileFolder, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	return filepath.Join(testFileFolder, "test")
}

func getExecutable(t *testing.T) string {
	// get os dependend file name
	filename := "fpcalc"
	switch runtime.GOOS {
	case "windows":
		filename = "win-" + filename + ".exe"
	case "darwin":
		filename = "macos-" + filename
	default:
		filename = "linux-" + filename
	}

	return filepath.Join(getTestFolderPath(t), "binary", filename)
}

func getDefaultChromaprint(t *testing.T) *Chromaprint {
	executable := getExecutable(t)

	result := Chromaprint{
		options: builder{
			filePath:             executable,
			sampleRateInHz:       -1,
			channels:             -1,
			maxFingerPrintLength: -1,
			chunkSizeInSeconds:   -1,
			overlap:              false,
			algorithm:            -1,
		},
	}
	return &result
}

func TestChromaprint_getArgs(t *testing.T) {
	tests := []struct {
		name string
		c    *Chromaprint
		want []string
	}{
		{
			name: "default",
			c:    getDefaultChromaprint(t),
			want: make([]string, 0),
		}, {
			name: "default",
			c: &Chromaprint{
				options: builder{
					filePath:             "test",
					sampleRateInHz:       44100,
					channels:             2,
					maxFingerPrintLength: 120,
					chunkSizeInSeconds:   5,
					overlap:              true,
					algorithm:            1,
				},
			},
			want: []string{"-rate", "44100", "-channels", "2", "-length", "120", "-chunk", "5", "-algorithm", "1", "-overlap"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.getArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chromaprint.getArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
