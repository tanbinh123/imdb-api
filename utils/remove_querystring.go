package utils

import "strings"

func RemoveQueryString(input string) string {
	return strings.Split(input, "?")[0]
}
