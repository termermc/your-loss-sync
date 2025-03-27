package util

import (
	"errors"
	"github.com/termermc/your-loss-sync/lang"
	"strconv"
	"strings"
)

// FormatBytes formats bytes into a human-readable string.
func FormatBytes(bytes int64) string {
	const (
		KiB = 1024
		MiB = KiB * 1024
		GiB = MiB * 1024
		TiB = GiB * 1024
	)

	if bytes < KiB {
		return strconv.FormatInt(bytes, 10) + " B"
	} else if bytes < MiB {
		value := float64(bytes) / float64(KiB)
		return strconv.FormatFloat(value, 'f', 1, 64) + " KiB"
	} else if bytes < GiB {
		value := float64(bytes) / float64(MiB)
		return strconv.FormatFloat(value, 'f', 1, 64) + " MiB"
	} else if bytes < TiB {
		value := float64(bytes) / float64(GiB)
		return strconv.FormatFloat(value, 'f', 1, 64) + " GiB"
	} else {
		value := float64(bytes) / float64(TiB)
		return strconv.FormatFloat(value, 'f', 1, 64) + " TiB"
	}
}

// ParseBytes parses a human-readable string into bytes.
// The string must be in the format `\s*\d+\s*(kb|kib|mb|mib|gb|gib|tb|tib)\s*`.
func ParseBytes(str string, locale lang.Locale) (int64, error) {
	str = strings.TrimSpace(str)

	// Collect number
	numStr := ""
	for i, r := range str {
		if i == 0 && r == '-' {
			return 0, errors.New(locale.Tr("parse.error.number-cannot-be-negative"))
		}

		if r == '．' || r == '。' {
			r = '.'
		}

		if (r >= '0' && r <= '9') || r == '.' {
			numStr += string(r)
		} else if r == ',' {
			return 0, errors.New(locale.Tr("parse.error.fractional-numbers-not-allowed"))
		} else {
			break
		}
	}
	if len(numStr) > 1 && numStr[0] == '.' && numStr[1] >= '0' && numStr[1] <= '9' {
		numStr = "0" + numStr
	}

	// Get suffix
	suffix := strings.ToLower(strings.TrimSpace(str[len(numStr):]))

	if suffix == "" {
		return 0, errors.New(locale.Tr("parse.error.missing-unit-suffix"))
	}

	if suffix == "b" {
		return strconv.ParseInt(numStr, 10, 64)
	}

	var mul int64
	switch suffix {
	case "kb":
		mul = 1000
		break
	case "kib":
		mul = 1024
		break
	case "mb":
		mul = 1000 * 1000
		break
	case "mib":
		mul = 1024 * 1024
		break
	case "gb":
		mul = 1000 * 1000 * 1000
		break
	case "gib":
		mul = 1024 * 1024 * 1024
		break
	case "tb":
		mul = 1000 * 1000 * 1000 * 1000
		break
	case "tib":
		mul = 1024 * 1024 * 1024 * 1024
		break
	default:
		return 0, errors.New(locale.Tr("parse.error.unknown-unit-x", suffix))
	}

	// Parse float
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, err
	}

	return int64(num * float64(mul)), nil
}
