package main

import (
	"strings"
)

func ContainsRune(s string, r rune) bool {
	i := strings.IndexRune(s, r)
	return i >= 0
}
