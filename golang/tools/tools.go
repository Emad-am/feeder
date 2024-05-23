package tools

import (
	"strconv"
	"strings"
)

func Trim(s string, chars []string) string {
	for i := 0; i < len(chars); i++ {
		s = strings.ReplaceAll(s, chars[i], "")
	}
	return s
}

func ToInt(s string) int64 {
	num, _ := strconv.ParseInt(s, 10, 64)
	return num
}
