// Copyright 2020 thinkgos (thinkgo@aliyun.com).  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package meter storage metering
package meter

import (
	"bytes"
	"math"
	"testing"
)

func issFloatEqual(x, y float64) bool {
	delta := math.Abs(x - y)
	mean := math.Abs(x+y) / 2.0
	return delta/mean < 0.000001
}

func TestByteSize(t *testing.T) {
	if got, want := ByteSize(100).Bytes(), uint64(100); got != want {
		t.Errorf("Bytes() = %v, want %v", got, want)
	}
	if got, want := ByteSize(1024*12).KBytes(), 12.0; !issFloatEqual(got, want) {
		t.Errorf("KBytes() = %v, want %v", got, want)
	}
	if got, want := ByteSize(1024*1024*12).MBytes(), 12.0; !issFloatEqual(got, want) {
		t.Errorf("MBytes() = %v, want %v", got, want)
	}
	if got, want := ByteSize(1024*1024*1024*12).GBytes(), 12.0; !issFloatEqual(got, want) {
		t.Errorf("GBytes() = %v, want %v", got, want)
	}
	if got, want := ByteSize(1024*1024*1024*1024*12).TBytes(), 12.0; !issFloatEqual(got, want) {
		t.Errorf("TBytes() = %v, want %v", got, want)
	}
	if got, want := ByteSize(1024*1024*1024*1024*1024*12).PBytes(), 12.0; !issFloatEqual(got, want) {
		t.Errorf("PBytes() = %v, want %v", got, want)
	}
	if got, want := ByteSize(1024*1024*1024*1024*1024*1024*13).EBytes(), 13.0; !issFloatEqual(got, want) {
		t.Errorf("EBytes() = %v, want %v", got, want)
	}
	got, err := ByteSize(5).MarshalText()
	if err != nil {
		t.Errorf("MarshalText() Err = %v, want %v", got, false)
	}
	if want := []byte("5B"); !bytes.Equal(got, want) {
		t.Errorf("MarshalText() = %v, want %v", got, want)
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
		{"5B", args{"5B"}, 5, false},
		{"B", args{"100.0B"}, 100, false},
		{"KB", args{"100.0KB"}, 1024 * 100, false},
		{"MB", args{"100.0MB"}, 1024 * 1024 * 100, false},
		{"GB", args{"100.0GB"}, 1024 * 1024 * 1024 * 100, false},
		{"TB", args{"100.0TB"}, 1024 * 1024 * 1024 * 1024 * 100, false},
		{"PB", args{"100.0PB"}, 1024 * 1024 * 1024 * 1024 * 1024 * 100, false},
		{"EB", args{"10.0EB"}, 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 10, false},
		{"no number", args{"b"}, 0, true},
		{"invalid unit", args{"101.0cb"}, 0, true},
		{"invalid number", args{"10.0.1B"}, 0, true},
		{"value out of range", args{"1000.0EB"}, math.MaxUint64, true},
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

func TestHumanSize(t *testing.T) {
	type args struct {
		bytes uint64
	}
	tests := []struct {
		name  string
		args  args
		wantS string
	}{
		{"5B", args{5}, "5B"},
		{"B", args{100}, "100.00B"},
		{"KB", args{1024 * 100}, "100.00KB"},
		{"MB", args{1024 * 1024 * 100}, "100.00MB"},
		{"GB", args{1024 * 1024 * 1024 * 100}, "100.00GB"},
		{"TB", args{1024 * 1024 * 1024 * 1024 * 100}, "100.00TB"},
		{"PB", args{1024 * 1024 * 1024 * 1024 * 1024 * 100}, "100.00PB"},
		{"EB", args{1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1}, "1.00EB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := HumanSize(tt.args.bytes); gotS != tt.wantS {
				t.Errorf("HumanSize() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
