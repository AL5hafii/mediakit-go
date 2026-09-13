package ffmpegArgs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Output adds the output file to the Args list and ensures the output directory exists
func (a *ARGS) Output(file string) error {
	file = strings.TrimSpace(file)
	if file == "" {
		return fmt.Errorf("output file cannot be empty")
	}

	err := os.MkdirAll(filepath.Dir(file), 0755)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	*a = append(*a, file)
	return nil
}
