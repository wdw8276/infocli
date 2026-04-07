package utils

import (
	"strings"
)

// NOTE: true or false
func HasSubStr(s string, sub string) bool {
	return strings.Contains(s, sub)
}

// NOTE: -1 not find substr, >0 return first match substr
func IndexOfSubStr(s string, sub string) int {
	return strings.Index(s, sub)
}

// NOTE: s and char must not "" and has char, else return space []
func Split(s string, char string) []string {
	slice := []string{}
	if len(s) > 0 && len(char) > 0 && IndexOfSubStr(s, char) >= 0 {
		slice = strings.Split(s, char)
	}
	return slice
}

func Join(s []string, char string) string {
	return strings.Join(s, char)
}
