package util

import (
	"strings"
)

func StringSliceToString(in []string) string {
	res := ""
	for k, v := range in {
		if k != 0 {
			res += ", "
		}
		res += v
	}
	return res
}

func SeparateComma(in string) []string {
	result := strings.Split(in, ",")
	for i := 0; i < len(result); i++ {
		result[i] = strings.TrimSpace(result[i])
	}

	return result
}

func MustGetString(s string, ok bool) string {
	if !ok {
		return ""
	}
	return s
}
