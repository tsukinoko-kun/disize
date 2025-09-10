package disize

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Size int

const (
	Bit = 1
	// Byte
	B = 8
	// Kilobyte
	KB = 1000 * B
	// Megabyte
	MB = 1000 * KB
	// Gigabyte
	GB = 1000 * MB
	// Terabyte
	TB = 1000 * GB
	// Petabyte
	PB = 1000 * TB
)

const (
	// Kilobit
	Kb = 1000 * B
	// Megabit
	Mb = 1000 * Kb
	// Gigabit
	Gb = 1000 * Mb
	// Terabit
	Tb = 1000 * Gb
	// Petabit
	Pb = 1000 * Tb
)

const (
	// Kibibyte
	KiB = 1024 * B
	// Mebibyte
	MiB = 1024 * KiB
	// Gibibyte
	GiB = 1024 * MiB
	// Tebibyte
	TiB = 1024 * GiB
	// Pebibyte
	PiB = 1024 * TiB
)

const (
	// Kibibit
	Kib = 1024 * Bit
	// Mebibit
	Mib = 1024 * Kib
	// Gibibit
	Gib = 1024 * Mib
	// Tebibit
	Tib = 1024 * Gib
	// Pebibit
	Pib = 1024 * Tib
)

func fmtMaxTwoDecimals(f float64) string {
	// Round to 2 decimal places
	r := math.Round(f*100) / 100
	// Print without trailing zeros
	return strconv.FormatFloat(r, 'f', -1, 64)
}

// String returns a human readable string representation of the size.
func (s Size) String() string {
	switch {
	case s >= PiB:
		return fmt.Sprintf("%s PiB", fmtMaxTwoDecimals(float64(s)/float64(PiB)))
	case s >= TiB:
		return fmt.Sprintf("%s TiB", fmtMaxTwoDecimals(float64(s)/float64(TiB)))
	case s >= GiB:
		return fmt.Sprintf("%s GiB", fmtMaxTwoDecimals(float64(s)/float64(GiB)))
	case s >= MiB:
		return fmt.Sprintf("%s MiB", fmtMaxTwoDecimals(float64(s)/float64(MiB)))
	case s >= KiB:
		return fmt.Sprintf("%s KiB", fmtMaxTwoDecimals(float64(s)/float64(KiB)))
	case s >= B:
		return fmt.Sprintf("%s B", fmtMaxTwoDecimals(float64(s)/float64(B)))
	default:
		return fmt.Sprintf("%d b", s)
	}
}

func (s Size) Bytes() int {
	return int(s / B)
}

// ParseSize parses a human readable string representation of a size.
func ParseSize(s string) (Size, error) {
	if len(s) == 0 {
		return 0, errors.New("empty size string")
	}
	s = strings.TrimSpace(s)
	switch s[len(s)-1] {
	case 'b':
		if len(s) < 2 {
			return 0, errors.New("unidentifiable size string")
		}
	case 'B':
		if len(s) < 2 {
			return 0, errors.New("unidentifiable size string")
		}
		switch s[len(s)-2] {
		case 'i':
			if len(s) < 3 {
				return 0, errors.New("unidentifiable size string")
			}
			switch s[len(s)-3] {
			case 'K', 'k':
				return parseSize(s[:len(s)-3], KiB)
			case 'M', 'm':
				return parseSize(s[:len(s)-3], MiB)
			case 'G', 'g':
				return parseSize(s[:len(s)-3], GiB)
			case 'T', 't':
				return parseSize(s[:len(s)-3], TiB)
			case 'P', 'p':
				return parseSize(s[:len(s)-3], PiB)
			default:
				return 0, fmt.Errorf("unknown unit %q", s[len(s)-3])
			}
		case 'K', 'k':
			return parseSize(s[:len(s)-2], KB)
		case 'M', 'm':
			return parseSize(s[:len(s)-2], MB)
		case 'G', 'g':
			return parseSize(s[:len(s)-2], GB)
		case 'T', 't':
			return parseSize(s[:len(s)-2], TB)
		case 'P', 'p':
			return parseSize(s[:len(s)-2], PB)
		default:
			if f, err := strconv.ParseFloat(strings.TrimSpace(s[:len(s)-1]), 64); err == nil {
				return Size(f * float64(B)), nil
			}
		}
	case 'K', 'k':
		return parseSize(s[:len(s)-1], KB)
	case 'M', 'm':
		return parseSize(s[:len(s)-1], MB)
	case 'G', 'g':
		return parseSize(s[:len(s)-1], GB)
	case 'T', 't':
		return parseSize(s[:len(s)-1], TB)
	case 'P', 'p':
		return parseSize(s[:len(s)-1], PB)
	}

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return Size(f * float64(B)), nil
	}

	return 0, fmt.Errorf("unidentifiable size string %q", s)
}

func parseSize(s string, base Size) (Size, error) {
	f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0, err
	}
	return Size(f * float64(base)), nil
}
