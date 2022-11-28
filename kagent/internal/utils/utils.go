package utils

import (
	"strings"
)

func BuildStringfunc(strs ...string) string {
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}