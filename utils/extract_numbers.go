package utils

import (
	"regexp"
	"strings"
)

func ExtractNumbers(input string) string {
	regex := regexp.MustCompile("[0-9]+")
	return strings.Join(regex.FindAllString(input, -1), "")
}
