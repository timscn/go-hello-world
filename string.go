package main

import (
	"strings"
)

func hasPrefix(str string) bool {
	return strings.HasPrefix(str, "Prefix")
}

func hasSufix(str string) bool {
	return strings.HasSuffix(str, "fix")
}
