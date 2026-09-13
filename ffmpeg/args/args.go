package ffmpegArgs

import "strings"

type ARGS []string

// Flag adds a flag and its corresponding value to the Args list
func (a *ARGS) AppendFlag(f string) {
	f = strings.TrimSpace(f)
	if f != "" {
		if !strings.Contains(f, "-") {
			f = "-" + f
		}
		*a = append(*a, f)
	}
}

// AppendValue adds a value to the Args list
func (a *ARGS) AppendValue(v string) {
	v = strings.TrimSpace(v)
	if v != "" {
		*a = append(*a, v)
	}
}

func (a *ARGS) Append(flag string, value string) {
	flag = strings.TrimSpace(flag)
	value = strings.TrimSpace(value)
	if flag != "" && value != "" {
		a.AppendFlag(flag)
		a.AppendValue(value)
	}
}
