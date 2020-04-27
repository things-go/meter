package meter

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var opt = cmp.Comparer(func(x, y float64) bool {
	delta := math.Abs(x - y)
	mean := math.Abs(x+y) / 2.0
	return delta/mean < 0.000001
})

func TestByteSize(t *testing.T) {
	if got, want := ByteSize(100).Bytes(), uint64(100); got != want {
		t.Errorf("Bytes() = %v, want %v", got, want)
	}
	if got, want := ByteSize(1024*12).KBytes(), 12.0; !cmp.Equal(got, want, opt) {
		t.Errorf("KBytes() = %v, want %v", got, want)
	}
	if got, want := ByteSize(1024*1024*12).MBytes(), 12.0; !cmp.Equal(got, want, opt) {
		t.Errorf("MBytes() = %v, want %v", got, want)
	}
	if got, want := ByteSize(1024*1024*1024*12).GBytes(), 12.0; !cmp.Equal(got, want, opt) {
		t.Errorf("GBytes() = %v, want %v", got, want)
	}
	if got, want := ByteSize(1024*1024*1024*1024*12).TBytes(), 12.0; !cmp.Equal(got, want, opt) {
		t.Errorf("TBytes() = %v, want %v", got, want)
	}
	if got, want := ByteSize(1024*1024*1024*1024*1024*12).PBytes(), 12.0; !cmp.Equal(got, want, opt) {
		t.Errorf("PBytes() = %v, want %v", got, want)
	}
	if got, want := ByteSize(1024*1024*1024*1024*1024*1024*12).EBytes(), 12.0; !cmp.Equal(got, want, opt) {
		t.Errorf("EBytes() = %v, want %v", got, want)
	}
}

func TestParseBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBytes(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
