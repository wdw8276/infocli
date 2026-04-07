package utils

import (
	"strconv"
)

func IntToStr(n int) string {
	return strconv.Itoa(n)
}

func StrToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

func StrToFloat(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return n
}

func IsNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
