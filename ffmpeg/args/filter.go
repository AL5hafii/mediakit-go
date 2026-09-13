package ffmpegArgs

import "strings"

// Filter appends the filter parameters to the Args
func (a *ARGS) Filter(flag string, graph ...string) {
	flag = strings.TrimSpace(flag)
	filters := make([]string, 0)

	for _, filter := range graph {
		filter = strings.TrimSpace(filter)
		if filter != "" {
			filters = append(filters, filter)
		}
	}

	if flag != "" && len(filters) > 0 {
		filter := strings.Join(filters, ";")
		switch flag {
		case "af":
			a.Append(flag, filter)
		case "vf":
			a.Append(flag, filter)
		case "filter_complex":
			a.Append(flag, filter)
		}
	}
}
