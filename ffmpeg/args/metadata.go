package ffmpegArgs

import (
	"fmt"
	"strings"
)

// METADATA represents a map of metadata key-value pairs for ffmpeg
type METADATA map[string]string

// Metadata appends the metadata parameters to the Args
func (a *ARGS) Metadata(m METADATA) {
	args := make(ARGS, 0)
	for key, value := range m {
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if key == "" || value == "" {
			continue
		}

		meta := fmt.Sprintf("%s=%s", key, value)
		args.Append("metadata", meta)
	}

	if len(args) > 0 {
		// Prepend "-metadata -1" to clear existing metadata before adding new metadata
		*a = append(ARGS{"-metadata", "-1"}, args...)
	}
}
