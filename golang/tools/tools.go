package tools

import (
	"encoding/json"
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

func StructToString(s interface{}) string {
	jsonVal, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(jsonVal)
}
