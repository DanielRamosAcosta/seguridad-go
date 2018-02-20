package utils

import (
	"strings"
)

/**
 * Cleans the given string.
 * Example: "-this IS NOT cleanñ-" => "THISISNOTCLEAN"
 */
func CleanString(str string) string {
	var clean string

	for _, c := range strings.ToUpper(str) {
		if 65 <= c && c <= 90 {
			clean = clean + string(c)
		}
	}

	return clean
}
