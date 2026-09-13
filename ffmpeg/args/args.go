package ffmpegArgs

import (
	"strings"
)

type Args []string

func (a *Args) Flag(flag string) {
	flag = strings.TrimSpace(flag)
	if flag != "" {
		*a = append(*a, flag)
	}
}

func (a *Args) Arg(flag string, value string) {
	flag = strings.TrimSpace(flag)
	value = strings.TrimSpace(value)
	if flag != "" && value != "" {
		*a = append(*a, flag, value)
	}
}
