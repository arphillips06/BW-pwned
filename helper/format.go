package helper

import (
	"regexp"
	"strings"
)

var ansi = regexp.MustCompile("\033\\[[0-9;]*m")

func StripAnsi(s string) string {
	return ansi.ReplaceAllString(s, "")
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func PadANSI(s string, width int) string {
	visible := StripAnsi(s)
	padding := max(width-len(visible), 0)
	return s + strings.Repeat(" ", padding)
}

const (
	Red    = "\033[1;31m"
	Green  = "\033[1;32m"
	Yellow = "\033[1;33m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
	Reset  = "\033[0m"
)
