package prct02

import (
	"bytes"
	"strings"
)

// CleanString cleans the given string.
// Example: "-this IS NOT cleanñ-" => "THISISNOTCLEAN"
func CleanString(str string) string {
	var buffer bytes.Buffer

	for _, c := range strings.ToUpper(str) {
		if 65 <= c && c <= 90 {
			buffer.WriteRune(c)
		}
	}

	return buffer.String()
}
