package util

import (
	"strings"
	"unicode"
)

func CleanString(s string) string {
	if s == "" {
		return s
	}

	s = strings.ReplaceAll(s, "\r", "")
	parts := strings.Fields(s)
	return strings.Join(parts, " ")
}

func CleanUpperString(s string) string {
	return strings.ToUpper(CleanString(s))
}

func RemoveCommentLines(s string) string {
	var out []string
	for _, line := range strings.Split(s, "\n") {
		var outLine []rune
		dashBefore := false
		for _, r := range line {
			if r == '-' {
				if dashBefore {
					outLine = outLine[:len(outLine)-1]
					break
				}
				dashBefore = true
			}
			outLine = append(outLine, r)
		}
		out = append(out, string(outLine))
	}
	return strings.Join(out, "\n")
}

func ShellSplit(s string) []string {
	var out []string
	var current []rune
	inQuote := false
	for _, r := range s {
		if !inQuote && unicode.IsSpace(r) {
			if len(current) > 0 {
				out = append(out, string(current))
				current = []rune{}
			}
			continue
		}
		if r == '"' || r == '\'' {
			inQuote = !inQuote
			continue
		}

		current = append(current, r)
	}
	if len(current) > 0 {
		out = append(out, string(current))
	}
	return out
}

func CleanSplit(s string, sep string) []string {
	var out []string
	for _, part := range strings.Split(s, sep) {
		cleaned := CleanString(part)
		if cleaned != "" {
			out = append(out, cleaned)
		}
	}
	return out
}
