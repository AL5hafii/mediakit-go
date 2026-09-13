package ffmpegArgs

import (
	"fmt"
	"strings"
)

// Maps adds the provided map values to the Args list with the "-map" flag
func (a *ARGS) Maps(ms ...string) {
	for _, m := range ms {
		m = strings.TrimSpace(m)
		if m != "" {
			// If the map does not contain a colon, wrap it in square brackets
			// This is to ensure that the map is treated as a stream specifier
			if !strings.Contains(m, ":") {
				m = fmt.Sprintf("[%s]", m)
			}

			*a = append(*a, "map", m)
		}
	}
}
