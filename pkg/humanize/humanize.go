package humanize

import (
	"fmt"
	"math"
)

// SI returns a humanized string for a given value using the SI (International System of Units) prefix.
func SI(value float64) string {
	if value < 0 {
		return "-"
	}

	if value == 0 {
		return "0"
	}

	sign := ""
	abs := math.Abs(value)

	format := func(n float64, suffix string) string {
		if n == math.Trunc(n) {
			return fmt.Sprintf("%s%.0f%s", sign, n, suffix)
		}
		return fmt.Sprintf("%s%.1f%s", sign, n, suffix)
	}

	switch {
	case abs >= 1000000000000:
		return format(value/1000000000000, "T")
	case abs >= 1000000000:
		return format(value/1000000000, "G")
	case abs >= 1000000:
		return format(value/1000000, "M")
	case abs >= 1000:
		return format(value/1000, "K")
	default:
		return format(value, "")
	}
}
