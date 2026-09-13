package ffmpegArgs

import (
	"fmt"
	"os/exec"
)

// Excute executes the ffmpeg command with the provided arguments
func (a *ARGS) Excute(bin string) error {
	// default to "ffmpeg" if no binary path is provided
	if bin == "" {
		bin = "ffmpeg"
	}

	// Execute the command
	err := exec.Command(bin, *a...).Run()
	if err != nil {
		return fmt.Errorf("failed to execute ffmpeg command: %w", err)
	}

	return nil
}
