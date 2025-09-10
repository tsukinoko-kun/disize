package disize_test

import (
	"testing"

	"github.com/tsukinoko-kun/disize"
)

func TestParseSize(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s       string
		want    disize.Size
		wantErr bool
	}{
		{
			name: "no unit",
			s:    "100",
			want: 100 * disize.B,
		},
		{
			name: "Byte",
			s:    "1B",
			want: 1 * disize.B,
		},
		{
			name: "Byte with space",
			s:    " 1 B ",
			want: 1 * disize.B,
		},
		{
			name: "K",
			s:    "1K",
			want: 1 * disize.KB,
		},
		{
			name: "K with space",
			s:    " 1 K ",
			want: 1 * disize.KB,
		},
		{
			name: "KB",
			s:    "11KB",
			want: 11 * disize.KB,
		},
		{
			name: "KB with space",
			s:    " 11 KB ",
			want: 11 * disize.KB,
		},
		{
			name: "PiB",
			s:    "42PiB",
			want: 42 * disize.PiB,
		},
		{
			name: "PiB with space",
			s:    " 42 PiB ",
			want: 42 * disize.PiB,
		},
		{
			name: "float",
			s:    "976.5625 KiB",
			want: 1 * disize.MB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := disize.ParseSize(tt.s)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ParseSize() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ParseSize() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("ParseSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSize_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		s    disize.Size
		want string
	}{
		{
			name: "B",
			s:    1 * disize.B,
			want: "1 B",
		},
		{
			name: "MB",
			s:    1 * disize.MB,
			want: "976.56 KiB",
		},
		{
			name: "MiB",
			s:    12 * disize.MiB,
			want: "12 MiB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.String()
			if got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
