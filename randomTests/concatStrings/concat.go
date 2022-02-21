package main

import (
	"fmt"
	"strings"
)

func sprintStyle(a, b string) string {
	return fmt.Sprintf("%s|%s", a, b)
}

func stringsStyle(a, b string) string {
	return strings.Join([]string{a, b}, "")
}

func stringsConcat(a, b string) string {
	return a + b
}
