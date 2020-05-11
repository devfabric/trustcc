package utils

import "strings"

func trimStringFromDot(s string) string {
	if idx := strings.LastIndexByte(s, '.'); idx != -1 {
		return s[idx+1:]
	}
	return s
}
