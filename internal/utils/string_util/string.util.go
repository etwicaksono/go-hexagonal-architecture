package string_util

import "strings"

func Explode(s, separator string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, separator)
}

func Implode(s []string, separator string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.Join(s, separator)
}

func ToLowerSlice(ids []string) []string {
	result := make([]string, len(ids))
	for i, id := range ids {
		result[i] = strings.ToLower(id)
	}
	return result
}
