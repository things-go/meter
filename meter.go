// Copyright 2020 thinkgos (thinkgo@aliyun.com).  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package meter storage metering
package meter

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Storage unit constants.
const (
	Byte  = 1
	KByte = Byte << 10
	MByte = KByte << 10
	GByte = MByte << 10
	TByte = GByte << 10
	PByte = TByte << 10
	EByte = PByte << 10
)

// ByteSize byte size
type ByteSize uint64

// Bytes to byte
func (b ByteSize) Bytes() uint64 {
	return uint64(b)
}

// KBytes to kilobyte
func (b ByteSize) KBytes() float64 {
	return float64(b/KByte) + float64(b%KByte)/float64(KByte)
}

// MBytes to megabyte
func (b ByteSize) MBytes() float64 {
	return float64(b/MByte) + float64(b%MByte)/float64(MByte)
}

// GBytes to gigabyte
func (b ByteSize) GBytes() float64 {
	return float64(b/GByte) + float64(b%GByte)/float64(GByte)
}

// TBytes to terabyte
func (b ByteSize) TBytes() float64 {
	return float64(b/TByte) + float64(b%TByte)/float64(TByte)
}

// PBytes to petabyte
func (b ByteSize) PBytes() float64 {
	return float64(b/PByte) + float64(b%PByte)/float64(PByte)
}

// EBytes to ebyte
func (b ByteSize) EBytes() float64 {
	return float64(b/EByte) + float64(b%EByte)/float64(EByte)
}

// String to string like xxB,xxKB,xxMB,xxGB,xxTB,xxPB,xxEB
func (b ByteSize) String() string {
	if b < 10 {
		return fmt.Sprintf("%dB", b)
	}
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	e := math.Floor(math.Log(float64(b)) / math.Log(1024))
	val := float64(b) / math.Pow(1024, e)
	return fmt.Sprintf("%0.1f%s", val, sizes[int(e)])
}

// HumanSize human readable string
func (b ByteSize) HumanSize() string {
	return b.String()
}

// MarshalText marshal to text
func (b ByteSize) MarshalText() ([]byte, error) {
	return []byte(b.String()), nil
}

// UnmarshalText unmarshal to ByteSize
func (b *ByteSize) UnmarshalText(t []byte) error {
	var i int

	// backup for error message
	t0 := t

	for ; i < len(t); i++ {
		if c := t[i]; !(('0' <= c && c <= '9') || c == '.') {
			if i == 0 {
				*b = 0
				return &strconv.NumError{
					Func: "UnmarshalText",
					Num:  string(t0),
					Err:  strconv.ErrSyntax,
				}
			}
			break
		}
	}

	val, err := strconv.ParseFloat(string(t[:i]), 10)
	if err != nil {
		return &strconv.NumError{
			Func: "UnmarshalText",
			Num:  string(t0),
			Err:  err,
		}
	}

	unit := uint64(Byte)
	unitStr := strings.ToLower(strings.TrimSpace(string(t[i:])))
	switch unitStr {
	case "", "b", "byte": // do nothing
	case "k", "kb", "kilo", "kilobyte", "kilobytes":
		unit = KByte
	case "m", "mb", "mega", "megabyte", "megabytes":
		unit = MByte
	case "g", "gb", "giga", "gigabyte", "gigabytes":
		unit = GByte
	case "t", "tb", "tera", "terabyte", "terabytes":
		unit = TByte
	case "p", "pb", "peta", "petabyte", "petabytes":
		unit = PByte
	case "e", "eb", "exa", "exabyte", "exabytes":
		unit = EByte
	default:
		*b = 0
		return &strconv.NumError{
			Func: "UnmarshalText",
			Num:  string(t0),
			Err:  strconv.ErrSyntax,
		}
	}
	if uint64(val) > math.MaxUint64/unit {
		*b = ByteSize(math.MaxUint64)
		return &strconv.NumError{
			Func: "UnmarshalText",
			Num:  string(t0),
			Err:  strconv.ErrRange,
		}
	}
	*b = ByteSize(uint64(val) * unit)
	return nil
}

// ParseBytes parse a human string to byte
func ParseBytes(s string) (uint64, error) {
	v := ByteSize(0)
	err := v.UnmarshalText([]byte(s))
	return v.Bytes(), err
}

// HumanSize human readable string
func HumanSize(bytes uint64) (s string) {
	return ByteSize(bytes).HumanSize()
}
